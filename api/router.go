package api

import (
	v1 "miniurl/api/v1"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()

	v1.NewRouter(s)
}
