package order

import (
	"github.com/gorilla/mux"
)

func Route(r *mux.Router, a *Api) {
	r.HandleFunc("", a.create).Methods("POST")
}
