package handlers

import (
	"net/http"
)

func (c *Context) Hello2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(c.Name))
}
