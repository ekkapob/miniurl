package handlers

import (
	"miniurl/api/models"
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

	rowDeleted, shortURL, err := c.URLService.DeleteURL(models.URL{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if rowDeleted != 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.URLService.DeleteCache(shortURL)
	w.WriteHeader(http.StatusOK)
}
