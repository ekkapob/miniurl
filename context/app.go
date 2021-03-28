package context

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

type App struct {
	DB *pg.DB
	RD *redis.Client
}
