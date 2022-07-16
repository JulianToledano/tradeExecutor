package server

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"tradeExecutor/order"
)

type App struct {
	r  *mux.Router
	db *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{
		r:  mux.NewRouter(),
		db: db,
	}
}

func (a *App) Run(addr string) {
	ordersApi := order.NewApi(a.db)
	order.Route(a.r.PathPrefix("/order").Subrouter(), ordersApi)

	srv := &http.Server{
		Handler:      a.r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}
