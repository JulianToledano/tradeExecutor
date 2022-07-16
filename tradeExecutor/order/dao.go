package order

import (
	"time"
	"tradeExecutor"
)

type dao struct {
	db    tradeExecutor.DataBase
	table string
}

func newDao(db tradeExecutor.DataBase) *dao {
	return &dao{
		db:    db,
		table: "orders",
	}
}

func (d dao) persist(o *Order) (err error) {
	return d.db.Insert(
		"INSERT INTO orders(size, price, symbol, buy, date) values(?,?,?,?,?)",
		o.Size, o.Price, o.Symbol, o.Buy, time.Now().Format("2006.01.02 15:04:05"))
}
