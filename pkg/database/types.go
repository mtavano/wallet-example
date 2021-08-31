package database

import "time"

// Tx is a user transaction
type Tx struct {
	Amount    int64
	CreatedAt time.Time
}

// Ledger is the basic user structure
type Ledger struct {
	TransactionsCurrency map[Currency][]*Tx
	BalancesCurrency     map[Currency]int64
}
