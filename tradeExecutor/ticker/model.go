package ticker

import "strconv"

type RawTicker struct {
	ID       uint64 `json:"u"`
	Symbol   string `json:"s"`
	BidPrice string `json:"b"`
	BidQty   string `json:"B"`
	AskPrice string `json:"a"`
	AskQty   string `json:"A"`
}

type Ticker struct {
	ID       uint64
	Symbol   string
	BidPrice float64
	BidQty   float64
	AskPrice float64
	AskQty   float64
}

func (raw RawTicker) Decode() (t *Ticker, err error) {
	t = &Ticker{}
	t.ID = raw.ID
	t.Symbol = raw.Symbol
	t.BidPrice, err = strconv.ParseFloat(raw.BidPrice, 32)
	if err != nil {
		return
	}
	t.BidQty, err = strconv.ParseFloat(raw.BidQty, 32)
	if err != nil {
		return
	}
	t.AskPrice, err = strconv.ParseFloat(raw.AskPrice, 32)
	if err != nil {
		return
	}
	t.AskQty, err = strconv.ParseFloat(raw.AskQty, 32)
	if err != nil {
		return
	}
	return
}
