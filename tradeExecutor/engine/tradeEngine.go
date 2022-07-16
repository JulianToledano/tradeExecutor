package engine

import (
	"encoding/json"
	"tradeExecutor/logger"
	"tradeExecutor/order"
	"tradeExecutor/ticker"
	"tradeExecutor/wss"
)

type tradeEngine struct {
	o         *order.Order
	webSocket wss.WebSocket
}

func NewTradeEngine(o *order.Order, kind string) *tradeEngine {
	return &tradeEngine{
		o:         o,
		webSocket: wss.Factory(kind, o),
	}
}

func (e tradeEngine) Execute() {
	logger.Infof(`[ENGINE] [tradeEngine] [Execute]`)
	c := make(chan []byte)
	go e.webSocket.ReadSocket(c)
	count := 0
	for {
		select {
		case msg, ok := <-c:
			if !ok {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [CHANNEL CLOSED]`)
				return
			}
			var t ticker.Ticker
			err := json.Unmarshal(msg, &t)
			if err != nil {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [%v]`, err)
				continue
			}
			count += 1
			if count == 50 {
				return
			}
		}
	}
}
