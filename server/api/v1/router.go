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

	v1.HandleFunc("/auth", ctx.Auth).Methods("POST")

	v1.HandleFunc("/urls",
		mw.CheckURL(mw.CheckBlacklist(ctx.CreateURL))).
		Methods("POST")

	v1.HandleFunc("/urls", mw.BasicAuth(ctx.GetURLs)).
		Methods("GET")
	v1.HandleFunc("/urls/{id}", mw.BasicAuth(ctx.DeleteURL)).
		Methods("DELETE")

	v1.HandleFunc("/urls/{shortURL}", ctx.GetURL).
		Methods("GET")

	v1.HandleFunc("/blacklist_urls", mw.BasicAuth(ctx.GetBlacklistURLs)).
		Methods("GET")
	v1.HandleFunc("/blacklist_urls",
		mw.BasicAuth(mw.CheckURL(ctx.CreateBlacklistURL))).
		Methods("POST")
	v1.HandleFunc("/blacklist_urls/{id}", mw.BasicAuth(ctx.DeleteBlacklistURL)).
		Methods("DELETE")
}
