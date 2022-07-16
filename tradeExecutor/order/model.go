package order

import (
	"errors"
	"net/http"
)

var (
	ErrSizeTooSmall = errors.New("order size is too small")
)

type Order struct {
	Size   float32 `json:"size"`
	Price  float32 `json:"price"`
	Symbol string  `json:"symbol"`
	Buy    bool    `json:"buy"`
}

func (o Order) Bind(r *http.Request) error {
	if o.Size <= 0 {
		return ErrSizeTooSmall
	}
	return nil
}

func (o Order) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
