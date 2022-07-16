package order

import (
	"database/sql"
	"fmt"
	"time"
)

type dao struct {
	db    *sql.DB
	table string
}

func newDao(db *sql.DB) *dao {
	return &dao{
		db:    db,
		table: "orders",
	}
}

func (d dao) persist(o *Order) (err error) {
	fmt.Sprint()
	_, err = d.db.Exec(
		"INSERT INTO orders(size, price, symbol, buy, date) values(?,?,?,?,?)",
		o.Size, o.Price, o.Symbol, o.Buy, time.Now().Format("2006.01.02 15:04:05"))
	return
}
