package order

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Api struct {
	next Service
}

func NewApi(db *sql.DB) *Api {
	return &Api{
		next: NewLogging(db),
	}
}

func (a Api) create(w http.ResponseWriter, r *http.Request) {
	var o *Order
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&o)
	if err != nil {
		return
	}
	err = a.next.create(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(nil)
}
