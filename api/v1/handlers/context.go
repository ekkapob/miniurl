package handlers

import "github.com/go-pg/pg/v10"

type Context struct {
	DB   *pg.DB
	Name string
}