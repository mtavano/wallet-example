package database

import "time"

// Tx is a user transaction
type Tx struct {
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

// Ledger is the basic user structure
type Ledger struct {
	TransactionsCurrency map[Currency][]*Tx `json:"transactions_currency"`
	BalancesCurrency     map[Currency]int64 `json:"balances_currency"`
}
