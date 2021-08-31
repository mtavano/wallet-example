package database

type CreateWithdrawInput struct {
	UserID string
	Amount uint64
}

func (st *Store) CreateWithdraw(input *CreateWithdrawInput) (int64, error) {
	return st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Amount:   -1 * int64(input.Amount),
		Currency: CurrencyUSDC,
	})
}

