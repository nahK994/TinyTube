package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	key := vars["id"]
	if id, err := strconv.Atoi(key); err != nil {
		return -1, err
	} else {
		return id, nil
	}
}
