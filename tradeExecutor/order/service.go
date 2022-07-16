package order

import (
	"tradeExecutor"
)

type Service interface {
	create(*Order) error
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
func (s service) create(o *Order) error {
	s.c <- o
	return s.d.persist(o)
}
