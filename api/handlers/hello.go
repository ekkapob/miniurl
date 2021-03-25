package handlers

import (
	"miniurl/internal"
	"net/http"
)

func Hello(ctx internal.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(ctx.Name))
	}
}
