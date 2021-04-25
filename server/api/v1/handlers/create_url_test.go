package handlers

import (
	"bytes"
	"errors"
	"miniurl/api/mid"
	"miniurl/pkg/base62"
	"miniurl/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func postCreateURL(url string) (*http.Request, *httptest.ResponseRecorder) {
	jsonStr := []byte(`{"url": "` + url + `"}`)
	req, _ := http.NewRequest("POST", "/api/v1/urls", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestInvalidURL(t *testing.T) {
	req, rr := postCreateURL("www.google.com")
	ctx := Context{URLService: &MockService{}}

	mw := &mid.Context{service.NewURLService(ctx.DB, ctx.RD)}
	handler := http.HandlerFunc(mw.CheckURL(ctx.CreateURL))
	handler.ServeHTTP(rr, req)

	expectedCode := http.StatusBadRequest
	if rr.Code != expectedCode {
		t.Errorf("expect status %v but got %v", expectedCode, rr.Code)
	}
}

func TestErrorInsertURL(t *testing.T) {
	url := "https://www.google.com"
	req, rr := postCreateURL(url)

	ctx := Context{
		URLService: &MockService{InsertURLError: errors.New("errro")},
	}

	handler := http.HandlerFunc(ctx.CreateURL)
	handler.ServeHTTP(rr, req)

	expectedCode := http.StatusInternalServerError
	if rr.Code != expectedCode {
		t.Errorf("expect status %v but got %v", expectedCode, rr.Code)
	}
}

func TestSuccessCreateURL(t *testing.T) {
	url := "https://www.google.com"
	req, rr := postCreateURL(url)

	mockCounter := 5
	ctx := Context{
		URLService: &MockService{Counter: mockCounter},
	}

	handler := http.HandlerFunc(ctx.CreateURL)
	handler.ServeHTTP(rr, req)

	var d struct {
		FullURL  string `json:"full_url"`
		ShortURL string `json:"short_url"`
	}
	DecodeJSON(rr.Body, &d)

	if d.FullURL != url {
		t.Errorf("expect full URL %v but got %v", url, d.FullURL)
	}
	shortURL := base62.Encode(uint64(mockCounter))
	if d.ShortURL != shortURL {
		t.Errorf("expect short URL %v but got %v", shortURL, d.ShortURL)
	}
}
