package wss

import (
	"tradeExecutor"
	"tradeExecutor/wss/binance"
	"tradeExecutor/wss/mock"
)

func Factory(wss string, symbol string) tradeExecutor.WebSocket {
	if wss == "binance" {
		return binance.NewBinanceWss(symbol)
	} else if wss == "mock" || wss == "test" {
		return mock.NewMockWss(nil)
	}
	return binance.NewBinanceWss(symbol)
}
