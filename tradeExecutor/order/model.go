package order

type Order struct {
	Size   float32 `json:"size"`
	Price  float32 `json:"price"`
	Symbol string  `json:"symbol"`
	Buy    bool    `json:"buy"`
}
