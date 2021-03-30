package mid

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Context) CheckBlacklist(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Req struct {
			URL string
		}
		var req Req
		b, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(b, &req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		urls, err := c.URLService.GetBlacklistURLs()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, v := range urls {
			if v == req.URL {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(
					struct {
						Error string `json:"error"`
					}{
						Error: "This URL is in blacklist.",
					})
				return
			}
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		next(w, r)
	}
}
