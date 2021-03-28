package mid

import (
	"miniurl/db"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Context) CheckCachedURL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		dbCtx := db.Context{
			DB: c.DB,
			RD: c.RD,
		}
		cached, err := dbCtx.GetCachedURL(shortURL)
		if err == nil {

			go func() {
				url, err := dbCtx.GetURLFromShortURL(shortURL)
				if err != nil {
					return
				}
				dbCtx.CacheURL(url)
				dbCtx.UpdateHit(url)
			}()

			http.Redirect(w, r, cached, http.StatusFound)
			return
		}
		next(w, r)
	}
}
