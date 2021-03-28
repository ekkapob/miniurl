package db

import (
	"errors"
	"miniurl/api/models"
	"strconv"
)

func (c *Context) InsertURL(url models.URL) error {
	_, err := c.DB.Model(&url).Insert()
	return err
}

func (c *Context) GetURLFromShortURL(shortURL string) (models.URL, error) {
	var url models.URL
	err := c.DB.Model(&url).
		Where(
			`short_url = ? AND
			now() <= created_at + expires_in_seconds * interval '1 second'`,
			shortURL,
		).
		Select()

	return url, err
}

func (c *Context) UpdateHit(url models.URL) error {
	c.mu.Lock()
	url.Hits++
	_, err := c.DB.Model(&url).WherePK().Update()
	c.mu.Unlock()
	return err
}

func (c *Context) GetURLs(options map[string]string) (
	urls []models.URL, total int, err error,
) {

	query := c.DB.Model(&urls)
	i, err := GetIntFromMap(options, "limit")
	if err == nil {
		query.Limit(i)
	}
	i, err = GetIntFromMap(options, "page")
	if err == nil {
		query.Offset(i)
	}

	if v, ok := options["orderBy"]; ok {
		if v == "expired_date" {
			query.OrderExpr(
				"created_at + expires_in_seconds * interval '1 second'" +
					" " +
					options["orderDirection"],
			)
		} else {
			query.Order(v + " " + options["orderDirection"])
		}
	}

	count, err := query.SelectAndCount()
	return urls, count, err
}

func GetIntFromMap(m map[string]string, key string) (int, error) {
	if v, ok := m[key]; ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i, nil
		}
	}
	return 0, errors.New("unable to find value")
}
