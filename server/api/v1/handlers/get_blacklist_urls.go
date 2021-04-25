package handlers

import (
	"encoding/json"
	"miniurl/api/models"
	"net/http"
)

func (c *Context) GetBlacklistURLs(w http.ResponseWriter, r *http.Request) {
	urls, err := c.URLService.GetBlacklistURLs()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		URLs []models.BlacklistURL `json:"urls"`
	}{urls})
}
