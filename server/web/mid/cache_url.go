package mid

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (c *Context) CheckCachedURL(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		cached, err := c.URLService.GetCachedURL(shortURL)
		if err == nil {

			go func() {
				url, err := c.URLService.GetURLFromShortURL(shortURL)
				if err != nil {
					return
				}
				c.URLService.CacheURL(url)
				c.URLService.UpdateHit(url)
			}()

			http.Redirect(w, r, cached, http.StatusFound)
			return
		}
		next(w, r)
	}
}
