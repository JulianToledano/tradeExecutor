package engine

import (
	"tradeExecutor/logger"
	"tradeExecutor/order"
)

type Executor interface {
	Execute()
}

type OrderEngine struct {
	c <-chan *order.Order
}

func NewTradeExecutor(c <-chan *order.Order) *OrderEngine {
	return &OrderEngine{
		c: c,
	}
}

func (e OrderEngine) Execute() {
	for {
		select {
		case o := <-e.c:
			logger.Infof("[ENGINE] [OrderEngine] [Execute]")
			tEngine := NewTradeEngine(o, "binance")
			go tEngine.Execute()
		}
	}
}
