package database

import "errors"

var ErrUserAlreadyExist = errors.New("user already exist")

func (st *Store) CreateUser(userid string) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	return st.createUser(userid)
}

func (st *Store) createUser(userid string) error {
	userID := UserID(userid)

	if _, ok := st.userLedgers[userID]; ok {
		return ErrUserAlreadyExist
	}

	transactionsCurrency := make(map[Currency][]*Tx)
	balancesCurrency := make(map[Currency]int64)

	for currency, _ := range supportedCurrencies {
		balancesCurrency[currency] = 0
		transactionsCurrency[currency] = make([]*Tx, 0)
	}

	st.userLedgers[userID] = &Ledger{
		TransactionsCurrency: transactionsCurrency,
		BalancesCurrency:     balancesCurrency,
	}

	return nil
}
