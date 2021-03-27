package web

import (
	"miniurl/web/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router) {
	r.HandleFunc("/{shortURL}", handlers.Redirect).Methods("GET")
}
