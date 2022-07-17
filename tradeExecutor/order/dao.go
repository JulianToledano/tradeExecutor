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
		"INSERT INTO orders(id, size, price, symbol, buy, date) values(?,?,?,?,?,?)",
		o.ID, o.Size, o.Price, o.Symbol, o.Buy, time.Now().Format("2006.01.02 15:04:05"))
}

func (d dao) list() (orders []Order, err error) {
	rows, err := d.db.List("SELECT * FROM orders")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var o Order
		err = rows.Scan(&o.ID, &o.Size, &o.Price, &o.Symbol, &o.Buy, &o.Date)
		if err != nil {
			continue
		}
		orders = append(orders, o)
	}
	return
}
