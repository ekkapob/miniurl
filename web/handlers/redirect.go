package handlers

import (
	"log"
	"miniurl/db"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	API_URL = "/api/v1/urls"
)

func (c *Context) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	dbCtx := &db.Context{
		DB: c.DB,
		RD: c.RD,
	}
	url, err := dbCtx.GetURLFromShortURL(shortURL)
	if err != nil {
		w.WriteHeader(http.StatusGone)
		return
	}
	err = dbCtx.UpdateHit(url)
	if err != nil {
		log.Println("error when update URL hit:", err)
	}

	dbCtx.CacheURL(url)
	http.Redirect(w, r, url.FullURL, http.StatusFound)
}
