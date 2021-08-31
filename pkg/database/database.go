package database

import (
	"sync"
	"time"
)

// Currency string alias
type Currency string
type UserID string

const (
	CurrencyUSDC Currency = "USDC"
	CurrencyBTC  Currency = "BTC"
	CurrencyETH  Currency = "ETH"
)

var supportedCurrencies = map[Currency]bool{
	CurrencyUSDC: true,
	CurrencyBTC:  true,
	CurrencyETH:  true,
}

// Store is the store struct
type Store struct {
	userLedgers map[UserID]*Ledger
	mu          sync.RWMutex

	now func() time.Time
}

func NewStore(timeFunc func() time.Time) *Store {
	return &Store{
		userLedgers: make(map[UserID]*Ledger),
		now:         timeFunc,
	}
}
