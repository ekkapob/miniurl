package db

import (
	"miniurl/api/models"
	"miniurl/pkg/utils"
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
	i, err := utils.GetIntFromMap(options, "limit")
	if err == nil {
		query.Limit(i)
	}
	i, err = utils.GetIntFromMap(options, "page")
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

func (c *Context) DeleteURL(url models.URL) (int, string, error) {
	var shortURLs []string
	r, err := c.DB.Model(&url).WherePK().Returning("short_url").Delete(&shortURLs)
	rowAffected := r.RowsAffected()

	var shortURL string
	if rowAffected > 0 {
		shortURL = shortURLs[0]
	}
	return rowAffected, shortURL, err
}
