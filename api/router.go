package api

import (
	v1 "miniurl/api/v1"
	"miniurl/api/v1/handlers"
	"miniurl/context"

	"github.com/gorilla/mux"
)

func NewRouter(appCtx context.App, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()

	ctx := handlers.Context{
		DB: appCtx.DB,
		RD: appCtx.RD,
	}

	v1.NewRouter(ctx, s)
}
