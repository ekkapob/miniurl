package main

import (
	"fmt"
	"log"
	"miniurl/api"
	"miniurl/context"
	"miniurl/web"
	"net/http"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func main() {
	// hostname := os.Getenv("HOSTNAME")
	port := os.Getenv("PORT")
	r := mux.NewRouter()

	ctx := context.App{
		RD: newRedis(),
		DB: newDB(),
	}

	api.NewRouter(ctx, r)
	web.NewRouter(ctx, r)

	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server is listening on", port)
	log.Fatal(srv.ListenAndServe())
}

func newRedis() *redis.Client {
	rd := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	return rd
}

func newDB() *pg.DB {
	opt, err := pg.ParseURL(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	return db
}
