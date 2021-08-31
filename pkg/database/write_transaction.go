package database

import (
	"errors"
)

var (
	ErrInsuficientFunds = errors.New("database: Store.WriteTransaction insufficient fund")
)

type writeTransactionInput struct {
	UserID   string
	Currency Currency
	Amount   int64
}

func (st *Store) writeTransaction(input *writeTransactionInput) (int64, error) {
	ledger, err := st.GetLedgerByUserID(input.UserID)
	if err != nil {
		return 0, nil
	}

	ledger, ok := st.userLedgers[UserID(input.UserID)]
	if !ok {
		return 0, ErrUserNotFound
	}

	prevBalance := ledger.BalancesCurrency[input.Currency]
	newBalance := prevBalance + input.Amount
	if newBalance < 0 {
		return 0, ErrInsuficientFunds
	}

	st.mu.Lock()
	defer st.mu.Unlock()

	ledger.BalancesCurrency[input.Currency] = newBalance
	ledger.TransactionsCurrency[input.Currency] = append(
		ledger.TransactionsCurrency[input.Currency],
		&Tx{
			Amount:    input.Amount,
			CreatedAt: st.now(),
		},
	)

	return newBalance, nil
}
