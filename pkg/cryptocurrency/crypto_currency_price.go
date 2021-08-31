package cryptocurrency

import (
	"strings"
	"sync"
)

type Wrapper struct {
	priceCurrency map[string]uint64
	mu            sync.RWMutex
}

func NewWrapper() *Wrapper {
	priceCurrency := map[string]uint64{
		"BTC": uint64(5000000), //50.000
		"ETH": uint64(400000),  //4.000
	}
	return &Wrapper{
		priceCurrency: priceCurrency,
	}
}

func (w *Wrapper) SetPrice(currency string, price uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	crypto := strings.ToUpper(currency)
	w.priceCurrency[crypto] = price
}

func (w *Wrapper) GetPrice(currency string) uint64 {
	w.mu.RLock()
	defer w.mu.RUnlock()

	crypto := strings.ToUpper(currency)
	return w.priceCurrency[crypto]
}
