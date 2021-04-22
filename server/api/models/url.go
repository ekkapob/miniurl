package models

import "time"

type URL struct {
	ID               int       `json:"id"`
	ShortURL         string    `json:"short_url"`
	FullURL          string    `json:"full_url"`
	Hits             int       `json:"hits"`
	CreatedAt        time.Time `json:"created_at"`
	ExpiresInSeconds int       `json:"expires_in_seconds"`
	LastModifiedAt   time.Time `json:"last_modified_at"`
}
