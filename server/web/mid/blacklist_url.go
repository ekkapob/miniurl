package mid

import (
	"net/http"

	"miniurl/api/models"

	"github.com/gorilla/mux"
)

func (c *Context) CheckBlacklist(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		url, err := c.URLService.GetURLFromShortURL(shortURL)
		if err != nil {
			w.WriteHeader(http.StatusGone)
			return
		}

		blacklistURL := models.BlacklistURL{URL: url.FullURL}
		err = c.URLService.FindBlacklistURL(&blacklistURL)
		if err == nil {
			w.WriteHeader(http.StatusGone)
			return
		}

		next(w, r)
	}
}
