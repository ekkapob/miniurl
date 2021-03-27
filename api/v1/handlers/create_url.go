package handlers

import (
	"encoding/json"
	"log"
	"miniurl/api/models"
	"miniurl/pkg/base62"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const URL_REGEXP = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

var expiresInSeconds = 604800

func init() {
	setURLExpiresInSeconds()
}

func setURLExpiresInSeconds() {
	i, err := strconv.Atoi(os.Getenv("URL_EXPIRE_SECONDS"))
	if err == nil && i >= 0 {
		expiresInSeconds = i
	}
}

func (c *Context) CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Req struct {
		URL string
	}

	var req Req
	err := DecodeReqJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isURLValid(req.URL) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			struct {
				Error string `json:"error"`
			}{
				Error: "URL is not valid. Please provide a full URL e.g. https://google.com",
			})
		return
	}

	// count, err := c.DB.Model(&models.URL{}).Count()
	// if err != nil {
	// 	log.Println("error when count urls:", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	counter, err := c.GetCounter()
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
	_, err = c.DB.Model(&url).Insert()
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

func isURLValid(url string) bool {
	if len(url) == 0 {
		return false
	}

	matched, err := regexp.Match(URL_REGEXP, []byte(url))
	if err != nil {
		return false
	}
	return matched
}
