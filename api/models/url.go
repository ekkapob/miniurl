package models

type URL struct {
	ShortURL         string `json:"short_url"`
	FullURL          string `json:"full_url"`
	Hits             int    `json:"hits"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
}
