package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Context) DeleteBlacklistURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	rowsAffected, err := c.URLService.DeleteBlacklistURL(id)
	if rowsAffected == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
