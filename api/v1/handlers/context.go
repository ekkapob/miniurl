package handlers

import (
	"sync"

	"github.com/go-pg/pg/v10"
)

type Context struct {
	DB *pg.DB
	mu sync.Mutex
}
