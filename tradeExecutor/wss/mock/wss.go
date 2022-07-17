package mock

import (
	"bytes"
	"encoding/gob"
	"tradeExecutor/logger"
	"tradeExecutor/ticker"
)

type mockWss struct {
	tickers []ticker.Ticker
}

func NewMockWss(tickers []ticker.Ticker) *mockWss {
	return &mockWss{
		tickers: tickers,
	}
}

func (mock mockWss) ReadSocket(ch chan<- []byte) {
	defer func() {
		logger.Warnf(`[WSS.MOCK] [MockWss] [ReadSocket] [CLOSING CHANNEL]`)
		close(ch)
	}()
	for _, t := range mock.tickers {
		var encTicker bytes.Buffer
		enc := gob.NewEncoder(&encTicker)
		err := enc.Encode(t)
		if err != nil {
			logger.Errorf(`[WSS.MOCK] [MockWss] [ReadSocket] [%v]`, err)
			return
		}
		ch <- encTicker.Bytes()
	}
}
