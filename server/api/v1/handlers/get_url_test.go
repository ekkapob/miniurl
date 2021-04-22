package handlers

import (
	"miniurl/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getURL(url string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/api/v1/url/a", nil)
	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestFoundGetURL(t *testing.T) {
	req, rr := getURL("x")
	url := models.URL{
		ShortURL: "x",
		FullURL:  "http:www.google.com",
	}
	ctx := Context{URLService: &MockService{GetURLData: url}}
	handler := http.HandlerFunc(ctx.GetURL)
	handler.ServeHTTP(rr, req)

	var d struct {
		FullURL  string `json:"full_url"`
		ShortURL string `json:"short_url"`
	}
	DecodeJSON(rr.Body, &d)
	if d.FullURL != url.FullURL {
		t.Errorf("expect full url %v but got %v", d.FullURL, url.FullURL)
	}
	if d.ShortURL != url.ShortURL {
		t.Errorf("expect short url %v but got %v", d.ShortURL, url.ShortURL)
	}
}
