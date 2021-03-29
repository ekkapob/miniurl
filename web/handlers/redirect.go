package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	API_URL = "/api/v1/urls"
)

func (c *Context) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	url, err := c.URLService.GetURLFromShortURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	err = c.URLService.UpdateHit(url)
	if err != nil {
		log.Println("error when update URL hit:", err)
	}

	c.URLService.CacheURL(url)
	http.Redirect(w, r, url.FullURL, http.StatusFound)
}
