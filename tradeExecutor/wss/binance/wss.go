package binance

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"tradeExecutor/logger"
)

const wssEndpoint = "wss://stream.binancefuture.com/ws/"

type BinanceWss struct {
	symbol  string
	baseUrl string
}

func NewBinanceWss(symbol string) *BinanceWss {
	return &BinanceWss{
		symbol:  symbol,
		baseUrl: wssEndpoint,
	}
}

func (bs BinanceWss) wssUrl(symbol string, stream string) string {
	return bs.baseUrl + strings.Join([]string{symbol, stream}, "@")
}

func (bs BinanceWss) bookTickerUrl(symbol string) string {
	return bs.wssUrl(symbol, "bookTicker")
}

func (bs BinanceWss) Dial(urlStr string) (*websocket.Conn, *http.Response, error) {
	return websocket.DefaultDialer.Dial(urlStr, nil)
}

func (bs BinanceWss) ReadSocket(ch chan<- []byte) {
	defer func() {
		logger.Warnf(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [CLOSING CHANNEL]`)
		close(ch)
	}()
	url := bs.bookTickerUrl(strings.ToLower(bs.symbol))
	c, _, err := bs.Dial(url)
	if err != nil {
		logger.Errorf(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [Dial] [%v]`, err)
		return
	}
	logger.Infof(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [%s]`, url)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			logger.Errorf(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [%v]`, err)
			continue
		}
		ch <- message
	}
}
