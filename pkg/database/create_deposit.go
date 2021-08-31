package database

type CreateDepositInput struct {
	UserID string
	Amount uint64
}

func (st *Store) CreateDeposit(input *CreateDepositInput) (int64, error) {
	return st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Amount:   int64(input.Amount),
		Currency: CurrencyUSDC,
	})
}
