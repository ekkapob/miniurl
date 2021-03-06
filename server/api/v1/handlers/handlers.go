package handlers

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader, v interface{}) error {
	d := json.NewDecoder(r)
	err := d.Decode(&v)
	if err != nil {
		return err
	}
	return nil
}
