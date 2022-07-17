package order

import (
	"errors"
	"github.com/google/uuid"
	"net/http"
)

var (
	ErrSizeTooSmall  = errors.New("order size is too small")
	ErrPriceTooSmall = errors.New("price is to small")
	ErrInvalidSymbol = errors.New("invalid symbol")
)

type Order struct {
	ID     string  `json:"id,omitempty"`
	Size   float64 `json:"size"`
	Price  float64 `json:"price"`
	Symbol string  `json:"symbol"`
	Buy    bool    `json:"buy"`
	Date   string  `json:"date,omitempty"`
}

func (o *Order) Bind(r *http.Request) error {
	if o.Size <= 0 {
		return ErrSizeTooSmall
	}
	if o.Price <= 0 {
		return ErrPriceTooSmall
	}
	if o.Symbol == "" {
		return ErrInvalidSymbol
	}
	o.ID = uuid.New().String()
	return nil
}

func (o Order) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
