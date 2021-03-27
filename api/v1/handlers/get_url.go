package handlers

import (
	"encoding/json"
	"miniurl/api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Context) GetURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	var url models.URL
	err := c.DB.Model(&url).
		Where(
			`short_url = ? AND
			now() <= created_at + expires_in_seconds * interval '1 second'`,
			shortURL,
		).
		Select()
	if err != nil {
		json.NewEncoder(w).Encode(
			struct {
				Error string `json:"error"`
			}{
				Error: "cannot find a full URL for the short URL",
			})
		return
	}

	json.NewEncoder(w).Encode(
		struct {
			ShortURL string `json:"short_url"`
			FullURL  string `json:"full_url"`
		}{
			ShortURL: url.ShortURL,
			FullURL:  url.FullURL,
		})
}
