package main

import (
	"fmt"
	"log"
	"miniurl/api"
	"miniurl/web"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
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
	hostname := os.Getenv("HOSTNAME")
	r := mux.NewRouter()

	api.NewRouter(r)
	web.NewRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server is listening on", hostname)
	log.Fatal(srv.ListenAndServe())
}
