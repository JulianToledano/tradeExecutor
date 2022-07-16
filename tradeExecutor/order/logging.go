package order

import (
	"database/sql"
	"fmt"
)

type logging struct {
	next Service
}

func NewLogging(db *sql.DB) *logging {
	return &logging{
		next: NewService(db),
	}
}

func (l logging) create(o *Order) (err error) {
	err = l.next.create(o)
	if err != nil {
		fmt.Println("ERROR creating order")
	}
	return
}
