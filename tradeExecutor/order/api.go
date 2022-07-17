package order

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"tradeExecutor"
)

type api struct {
	next Service
}

func NewApi(db tradeExecutor.DataBase, c chan<- *Order) *api {
	return &api{
		next: NewLogging(db, c),
	}
}

func (a api) create(w http.ResponseWriter, r *http.Request) {
	req := new(Order)
	if err := render.Bind(r, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := a.next.create(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (a api) list(w http.ResponseWriter, r *http.Request) {
	orders, err := a.next.list()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, orders)
}
