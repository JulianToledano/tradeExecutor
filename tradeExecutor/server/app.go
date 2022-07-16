package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"tradeExecutor"
	"tradeExecutor/logger"
	"tradeExecutor/order"
)

type App struct {
	r  *mux.Router
	db tradeExecutor.DataBase
}

func NewApp(db tradeExecutor.DataBase) *App {
	return &App{
		r:  mux.NewRouter(),
		db: db,
	}
}

func (a *App) Run(addr string, c chan<- *order.Order) {
	ordersApi := order.NewApi(a.db, c)
	order.Route(a.r.PathPrefix("/order").Subrouter(), ordersApi)

	srv := &http.Server{
		Handler:      a.r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Infof("[SERVER] [App] [Run] [%s]", addr)
	logger.Fatalf(srv.ListenAndServe().Error())
}
