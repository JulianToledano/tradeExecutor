package trade

import (
	"tradeExecutor/order"
	"tradeExecutor/ticker"
)

type Trade struct {
	Amount  float64
	Price   float64
	Buy     bool
	OrderId string
	Date    string
}

func NewTrade(t *ticker.Ticker, o *order.Order) (makeTrade bool, trade *Trade) {
	trade = &Trade{
		Buy:     o.Buy,
		OrderId: o.ID,
	}
	if o.Buy && t.AskPrice <= o.Price {
		if t.AskQty <= o.Size {
			trade.Amount = t.AskQty
		} else {
			trade.Amount = o.Size
		}
		trade.Price = t.AskPrice
		makeTrade = true
	} else if !o.Buy && t.BidPrice >= o.Price {
		if t.BidQty <= o.Size {
			trade.Amount = t.BidQty
		} else {
			trade.Amount = o.Size
		}
		trade.Price = t.BidPrice
		makeTrade = true
	}
	return
}
