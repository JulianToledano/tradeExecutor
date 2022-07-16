package binance

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"tradeExecutor/logger"
	"tradeExecutor/order"
)

const wssEndpoint = "wss://stream.binancefuture.com/ws/"

type BinanceWss struct {
	o       *order.Order
	baseUrl string
}

func NewBinanceWss(o *order.Order) *BinanceWss {
	return &BinanceWss{
		o:       o,
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
	url := bs.bookTickerUrl(strings.ToLower(bs.o.Symbol))
	c, _, err := bs.Dial(url)
	defer func() {
		logger.Warnf(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [CLOSING CHANNEL]`)
		close(ch)
	}()
	if err != nil {
		return
	}
	logger.Infof(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [%s]`, url)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			logger.Errorf(`[WSS.BINANCE] [BinanceWss] [ReadSocket] [%v]`, err)
			break
		}
		ch <- message
	}

}
