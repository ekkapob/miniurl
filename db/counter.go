package db

import "github.com/go-pg/pg/v10"

func (c *Context) GetCounter() (int, error) {
	var counter int
	c.mu.Lock()
	_, err := c.DB.QueryOne(
		pg.Scan(&counter),
		`SELECT nextval('url_counter')`,
	)
	c.mu.Unlock()
	return counter, err
}
