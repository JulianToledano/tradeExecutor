package trade

import (
	"time"
	"tradeExecutor"
)

type Dao struct {
	db    tradeExecutor.DataBase
	table string
}

func NewDao(db tradeExecutor.DataBase) *Dao {
	return &Dao{
		db:    db,
		table: "trades",
	}
}

func (d Dao) Persist(t *Trade) (err error) {
	return d.db.Insert(
		"INSERT INTO trades(amount, price, buy, orderId, date) values(?,?,?,?,?)",
		t.Amount, t.Price, t.Buy, t.OrderId, time.Now().Format("2006.01.02 15:04:05"))
}
