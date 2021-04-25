package handlers

import (
	"net/http"
)

func (c *Context) CreateBlacklistURL(w http.ResponseWriter, r *http.Request) {
	type Req struct {
		URL string
	}

	var req Req
	err := DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = c.URLService.InsertBlacklistURL(req.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
