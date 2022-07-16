package wss

import (
	"tradeExecutor/order"
	"tradeExecutor/wss/binance"
)

type WebSocket interface {
	ReadSocket(chan<- []byte)
}

func Factory(wss string, o *order.Order) WebSocket {
	if wss == "binance" {
		return binance.NewBinanceWss(o)
	}
	return binance.NewBinanceWss(o)
}
