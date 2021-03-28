package handlers

import (
	"miniurl/api/models"
	"miniurl/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Context) DeleteURL(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbCtx := db.Context{
		DB: c.DB,
		RD: c.RD,
	}
	rowDeleted, shortURL, err := dbCtx.DeleteURL(models.URL{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if rowDeleted != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dbCtx.DeleteCache(shortURL)
	w.WriteHeader(http.StatusOK)
}
