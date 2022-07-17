package engine

import (
	"tradeExecutor"
	"tradeExecutor/logger"
	"tradeExecutor/order"
)

type OrderEngine struct {
	wssName string
	c       <-chan *order.Order
	db      tradeExecutor.DataBase
}

func NewOrderEngine(wssName string, c <-chan *order.Order, db tradeExecutor.DataBase) *OrderEngine {
	return &OrderEngine{
		wssName: wssName,
		c:       c,
		db:      db,
	}
}

func (e OrderEngine) Execute() {
	for {
		select {
		case o, ok := <-e.c:
			if !ok {
				logger.Infof("[ENGINE] [OrderEngine] [Execute] [CHANNEL CLOSED]")
			}
			logger.Infof("[ENGINE] [OrderEngine] [Execute]")
			tEngine := NewTradeEngine(o, e.wssName, e.db)
			go tEngine.Execute()
		}
	}
}
