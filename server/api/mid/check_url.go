package mid

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"miniurl/pkg/url"
	"net/http"
)

func (c *Context) CheckURL(next http.HandlerFunc) http.HandlerFunc {
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

		if !url.IsValidURL(req.URL) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(
				struct {
					Error string `json:"error"`
				}{
					Error: "URL is not valid. Please provide a full URL e.g. https://google.com",
				})
			return
		}

		r.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		next(w, r)
	}
}
