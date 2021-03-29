package handlers

import (
	"miniurl/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getURLs(params string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/api/v1/urls"+params, nil)
	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetURLs(t *testing.T) {
	req, rr := getURLs("")
	urls := []models.URL{
		models.URL{
			ShortURL:         "a",
			FullURL:          "http://a.com",
			CreatedAt:        time.Now(),
			ExpiresInSeconds: 500,
		},
		models.URL{ShortURL: "b", FullURL: "http://b.com"},
		models.URL{ShortURL: "c", FullURL: "http://c.com"},
		models.URL{
			ShortURL:         "d",
			FullURL:          "http://d.com",
			CreatedAt:        time.Now(),
			ExpiresInSeconds: 1000,
		},
	}
	ctx := Context{URLService: &MockService{GetURLsData: urls}}
	handler := http.HandlerFunc(ctx.GetURLs)
	handler.ServeHTTP(rr, req)

	var d struct {
		URLs []models.URL `json:"urls"`
	}
	DecodeJSON(rr.Body, &d)
	if len(d.URLs) != len(urls) {
		t.Errorf(
			"expect to have %v items but got %v",
			len(d.URLs),
			len(urls),
		)
	}
}
