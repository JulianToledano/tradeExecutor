package engine

import (
	"testing"
	"tradeExecutor"
	"tradeExecutor/db"
	"tradeExecutor/order"
)

func TestEngine(t *testing.T) {
	tests := []struct {
		name string
		o    *order.Order
		c    chan *order.Order
		d    tradeExecutor.DataBase
	}{
		{
			name: "Test engine",
			o:    &order.Order{ID: "test-1", Size: 10, Price: 100, Symbol: "BTCUSDC", Buy: true},
			c:    make(chan *order.Order),
			d:    db.NewMockDb(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewOrderEngine("mock", tt.c, tt.d)
			go func() {
				tt.c <- tt.o
			}()
			e.Execute()
		})
	}
}
