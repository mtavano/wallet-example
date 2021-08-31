package database

import "errors"

var ErrUserNotFound = errors.New("database: user not found")

func (st *Store) GetLedgerByUserID(userID string) (*Ledger, error) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	userLedger, ok := st.userLedgers[UserID(userID)]
	if !ok {
		return nil, ErrUserNotFound
	}

	return userLedger, nil
}
