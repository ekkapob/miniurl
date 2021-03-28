package handlers

import (
	"encoding/json"
	"math"
	"miniurl/api/models"
	"miniurl/db"
	"net/http"
	"strconv"
)

func (c *Context) GetURLs(w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	orderBy := r.FormValue("orderBy")
	orderDirection := r.FormValue("orderDirection")

	dbCtx := db.Context{
		DB: c.DB,
	}
	queryOpts := map[string]string{
		"page":           page,
		"limit":          limit,
		"orderBy":        orderBy,
		"orderDirection": orderDirection,
	}
	urls, total, err := dbCtx.GetURLs(queryOpts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp struct {
		CurrentPage int          `json:"page,omitempty"`
		TotalPages  int          `json:"total_pages,omitempty"`
		URLs        []models.URL `json:"urls"`
	}

	resp.URLs = urls
	p, err := strconv.Atoi(page)
	if err == nil {
		resp.CurrentPage = p
	}
	l, err := strconv.Atoi(limit)
	if err == nil {
		resp.TotalPages = int(math.Ceil(float64(total) / float64(l)))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
