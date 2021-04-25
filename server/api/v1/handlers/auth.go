package handlers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

func (c *Context) Auth(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		Account  string
		Password string
	}
	var req Req
	err := DecodeJSON(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: check admin credentials from DB with bcrypt password
	if req.Account != os.Getenv("ADMIN_ACCOUNT") ||
		req.Password != os.Getenv("ADMIN_PASSWORD") {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	basicAuth := base64.StdEncoding.EncodeToString(
		[]byte(req.Account + ":" + req.Password),
	)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		struct {
			BasicAuth string `json:"basic_auth"`
		}{basicAuth})
}
