package order

import (
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, a *api) {
	r.HandleFunc("", a.create).Methods("POST")
	r.HandleFunc("", a.list).Methods("GET")
}
