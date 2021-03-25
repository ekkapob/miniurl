package handlers

import (
	"fmt"
	"miniurl/internal"
	"net/http"
)

func Hello(ctx internal.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("web...")
		w.Write([]byte(ctx.Name))
	}
}
