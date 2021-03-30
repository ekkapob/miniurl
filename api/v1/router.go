package v1

import (
	"miniurl/api/mid"
	"miniurl/api/v1/handlers"
	"miniurl/service"

	"github.com/gorilla/mux"
)

func NewRouter(ctx handlers.Context, r *mux.Router) {
	ctx.URLService = service.NewURLService(ctx.DB, ctx.RD)

	v1 := r.PathPrefix("/v1").Subrouter()

	mw := &mid.Context{service.NewURLService(ctx.DB, ctx.RD)}

	v1.HandleFunc("/urls", mw.CheckURL(mw.CheckBlacklist(ctx.CreateURL))).
		Methods("POST")
	v1.HandleFunc("/urls", ctx.GetURLs).
		Methods("GET")
	v1.HandleFunc("/urls/{id}", ctx.DeleteURL).
		Methods("DELETE")

	v1.HandleFunc("/urls/{shortURL}", ctx.GetURL).
		Methods("GET")

	v1.HandleFunc("/blacklist_urls", mw.CheckURL(ctx.CreateBlacklistURL)).
		Methods("POST")
}
