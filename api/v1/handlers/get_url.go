package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Context) GetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	url, err := c.URLService.GetURLFromShortURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			struct {
				Error string `json:"error"`
			}{
				Error: "cannot find a full URL for the short URL",
			})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		struct {
			ShortURL string `json:"short_url"`
			FullURL  string `json:"full_url"`
		}{
			ShortURL: url.ShortURL,
			FullURL:  url.FullURL,
		})
}
