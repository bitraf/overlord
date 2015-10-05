package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func queryParam(r *http.Request, name string) string {
	return r.URL.Query().Get(name)
}

func pathParam(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}

func parseInt(value string, def int) int {
	if len(value) == 0 {
		return def
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		return def
	}

	return val
}

func serveJson(w http.ResponseWriter, status int, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(result)
}
