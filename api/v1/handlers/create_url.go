package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"miniurl/api/models"
	"miniurl/pkg/base62"
	"miniurl/pkg/utils"
	"net/http"
	"os"
)

var expiresInSeconds = 604800

func init() {
	expiresInSeconds = utils.GetInt(
		os.Getenv("POSTGRES_URL_EXPIRE_SECONDS"),
		expiresInSeconds,
	)
}

func (c *Context) CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type Req struct {
		URL string
	}

	var req Req
	err := DecodeJSON(r.Body, &req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	counter, err := c.URLService.GetCounter()
	if err != nil {
		log.Println("error when get counter:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	shortURL := base62.Encode(uint64(counter))
	url := models.URL{
		ShortURL:         shortURL,
		FullURL:          req.URL,
		ExpiresInSeconds: expiresInSeconds,
	}

	err = c.URLService.InsertURL(url)
	if err != nil {
		log.Println("error when insert a url:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(
		struct {
			ShortURL string `json:"short_url"`
			FullURL  string `json:"full_url"`
		}{
			ShortURL: shortURL,
			FullURL:  req.URL,
		})
}
