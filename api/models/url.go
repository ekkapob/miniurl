package models

import "time"

type URL struct {
	ID               int
	ShortURL         string
	FullURL          string
	Hits             int
	ExpiresInSeconds int
	LastModifiedAt   time.Time
}
