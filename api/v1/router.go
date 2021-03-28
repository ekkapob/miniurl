package v1

import (
	"miniurl/api/v1/handlers"

	"github.com/gorilla/mux"
)

func NewRouter(ctx handlers.Context, r *mux.Router) {
	v1 := r.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/urls", ctx.CreateURL).
		Methods("POST")
	v1.HandleFunc("/urls", ctx.GetURLs).
		Methods("GET")
	v1.HandleFunc("/urls/{id}", ctx.DeleteURL).
		Methods("DELETE")

	v1.HandleFunc("/urls/{shortURL}", ctx.GetURL).
		Methods("GET")
}
