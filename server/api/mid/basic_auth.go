package mid

import (
	"encoding/base64"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (c *Context) BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		reg := regexp.MustCompile(`(?i)basic (.+)`)
		matches := reg.FindAllStringSubmatch(authHeader, 1)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token := matches[0][1]
		decoded, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		auth := strings.Split(string(decoded), ":")

		// TODO: check admin credentials from DB with bcrypt password
		if auth[0] != os.Getenv("ADMIN_ACCOUNT") ||
			auth[1] != os.Getenv("ADMIN_PASSWORD") {

			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

// when check password with bcrypt encoded
func isValidAdminPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(os.Getenv("ADMIN_PASSWORD")),
		[]byte(password),
	)
	return err == nil
}
