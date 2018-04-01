package code

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/data", Handler).Methods(http.MethodGet)

	return r
}