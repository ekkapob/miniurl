package web

import (
	"miniurl/context"
	"miniurl/service"
	"miniurl/web/handlers"
	"miniurl/web/mid"

	"github.com/gorilla/mux"
)

func NewRouter(appCtx context.App, r *mux.Router) {
	mw := &mid.Context{
		URLService: service.NewURLService(appCtx.DB, appCtx.RD),
	}

	ctx := &handlers.Context{
		URLService: service.NewURLService(appCtx.DB, appCtx.RD),
	}

	r.HandleFunc("/{shortURL}", mw.CheckCachedURL(ctx.Redirect)).
		Methods("GET")
}
