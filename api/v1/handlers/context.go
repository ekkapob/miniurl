package handlers

import (
	"miniurl/service"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
)

type Context struct {
	DB *pg.DB
	RD *redis.Client

	URLService service.URLService
}
