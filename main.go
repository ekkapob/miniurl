package main

import (
	"flag"
	"fmt"
	"log"
	"miniurl/api"
	"miniurl/internal"
	"miniurl/pkg/base62"
	"miniurl/web"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

const (
	serverAddr = "127.0.0.1"
)

type Counter struct {
	mu      sync.Mutex
	counter uint64
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func main() {
	port := flag.String("port", ":8000", "server port")
	flag.Parse()
	r := mux.NewRouter()

	fmt.Println(base62.Encode(123))

	ctx := internal.Context{
		Name: "hello",
	}

	web.NewRouter(ctx, r)
	api.NewRouter(ctx, r)

	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddr + *port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server is listening on", *port)
	log.Fatal(srv.ListenAndServe())
}
