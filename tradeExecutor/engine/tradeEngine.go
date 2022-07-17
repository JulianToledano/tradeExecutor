package engine

import (
	"encoding/json"
	"tradeExecutor"
	"tradeExecutor/logger"
	"tradeExecutor/order"
	"tradeExecutor/ticker"
	"tradeExecutor/trade"
	"tradeExecutor/wss"
)

type tradeEngine struct {
	o         *order.Order
	webSocket tradeExecutor.WebSocket
	db        tradeExecutor.DataBase
	tradeDao  *trade.Dao
}

func NewTradeEngine(o *order.Order, kind string, db tradeExecutor.DataBase) *tradeEngine {
	return &tradeEngine{
		o:         o,
		webSocket: wss.Factory(kind, o.Symbol),
		db:        db,
		tradeDao:  trade.NewDao(db),
	}
}

func (e tradeEngine) Execute() {
	logger.Infof(`[ENGINE] [tradeEngine] [Execute]`)
	c := make(chan []byte)
	go e.webSocket.ReadSocket(c)
	for {
		select {
		case msg, ok := <-c:
			if !ok {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [CHANNEL CLOSED]`)
				return
			}
			var raw ticker.RawTicker
			err := json.Unmarshal(msg, &raw)
			if err != nil {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [%v]`, err)
				continue
			}
			t, err := raw.Decode()
			if err != nil {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [DecodeTicker] [%v]`, err)
				continue
			}
			makeTrade, newTrade := trade.NewTrade(t, e.o)
			if makeTrade {
				err = e.tradeDao.Persist(newTrade)
				if err != nil {
					logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [InsertTrade] [%v]`, err)
					continue
				}
				e.o.Size -= newTrade.Amount
			}
			if e.o.Size <= 0 {
				logger.Errorf(`[ENGINE] [tradeEngine] [Execute] [FINISH] [%v]`, e.o.ID)
				return
			}
		}
	}
}
