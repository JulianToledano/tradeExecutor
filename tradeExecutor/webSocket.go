package tradeExecutor

type WebSocket interface {
	ReadSocket(chan<- []byte)
}
