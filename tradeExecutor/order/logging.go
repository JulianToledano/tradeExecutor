package order

import (
	"tradeExecutor"
	"tradeExecutor/logger"
)

type logging struct {
	next Service
}

func NewLogging(db tradeExecutor.DataBase, c chan<- *Order) *logging {
	return &logging{
		next: NewService(db, c),
	}
}

func (l logging) create(o *Order) (err error) {
	logger.Infof(`[ORDER] [Api] [create]`)
	err = l.next.create(o)
	if err != nil {
		logger.Errorf(`[ORDER] [Api] [create] [%v]`, err)
	}
	return
}

func (l logging) list() (orders []Order, err error) {
	logger.Infof(`[ORDER] [Api] [list]`)
	orders, err = l.next.list()
	if err != nil {
		logger.Errorf(`[ORDER] [Api] [list] [%v]`, err)
	}
	return
}
