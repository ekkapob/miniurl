package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	API_URL = "/api/v1/urls"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	url := fmt.Sprint(
		os.Getenv("HOSTNAME"),
		API_URL,
		"/",
		shortURL,
	)

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	var result map[string]interface{}
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusGone)
		return
	}

	json.NewDecoder(resp.Body).Decode(&result)
	fullURL := result["full_url"]
	if url, ok := fullURL.(string); ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	w.WriteHeader(http.StatusGone)
}
