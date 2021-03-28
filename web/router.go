package web

import (
	"miniurl/context"
	"miniurl/web/handlers"
	"miniurl/web/mid"

	"github.com/gorilla/mux"
)

func NewRouter(appCtx context.App, r *mux.Router) {
	mw := &mid.Context{
		DB: appCtx.DB,
		RD: appCtx.RD,
	}

	ctx := &handlers.Context{
		DB: appCtx.DB,
		RD: appCtx.RD,
	}

	r.HandleFunc("/{shortURL}", mw.CheckCachedURL(ctx.Redirect)).
		Methods("GET")
}
