package order

import (
	"tradeExecutor"
)

type Service interface {
	create(*Order) error
	list() ([]Order, error)
}

type service struct {
	c chan<- *Order
	d *dao
}

func NewService(db tradeExecutor.DataBase, c chan<- *Order) *service {
	return &service{
		c: c,
		d: newDao(db),
	}
}
func (s service) create(o *Order) (err error) {
	err = s.d.persist(o)
	if err == nil {
		s.c <- o
	}
	return
}

func (s service) list() ([]Order, error) {
	return s.d.list()
}
