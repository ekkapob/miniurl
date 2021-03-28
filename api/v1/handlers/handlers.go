package handlers

import (
	"encoding/json"
	"io"
)

func DecodeReqJSON(r io.Reader, v interface{}) error {
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	if err != nil {
		return err
	}
	return nil
}

// func (c *Context) GetCounter() (int, error) {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	var counter int
// 	_, err := c.DB.QueryOne(pg.Scan(&counter), `SELECT nextval('url_counter')`)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return counter, nil
// }
