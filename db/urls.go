package db

import (
	"miniurl/api/models"
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
