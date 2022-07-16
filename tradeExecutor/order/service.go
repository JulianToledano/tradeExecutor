package order

import "database/sql"

type Service interface {
	create(*Order) error
}

type service struct {
	d *dao
}

func NewService(db *sql.DB) *service {
	return &service{
		d: newDao(db),
	}
}
func (s service) create(o *Order) error {
	return s.d.persist(o)
}
