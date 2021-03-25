package api

import (
	"miniurl/api/handlers"
	"miniurl/internal"

	"github.com/gorilla/mux"
)

func NewRouter(ctx internal.Context, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()
	v1 := s.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/hello", handlers.Hello(ctx))
}
