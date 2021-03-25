package web

import (
	"miniurl/internal"
	"miniurl/web/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(ctx internal.Context, r *mux.Router) {
	r.HandleFunc("/", handlers.Hello(ctx))
}
