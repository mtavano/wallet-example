package database

type CreateTransferInput struct {
	SourceID      string
	DestinationID string
	Currency      Currency
	Amount        uint64
}

func (st *Store) CreateTransfer(input *CreateTransferInput) (int64, int64, error) {
	// check source id exist
	_, err := st.GetLedgerByUserID(input.SourceID)
	if err != nil {
		return 0, 0, err
	}

	// check destination id exist
	_, err = st.GetLedgerByUserID(input.DestinationID)
	if err != nil {
		return 0, 0, err
	}

	newSourceBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.SourceID,
		Amount:   -1 * int64(input.Amount),
		Currency: input.Currency,
	})
	if err != nil {
		return 0, 0, err
	}

	newDestinationBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.DestinationID,
		Amount:   int64(input.Amount),
		Currency: input.Currency,
	})
	if err != nil {
		return 0, 0, err
	}

	return newSourceBalance, newDestinationBalance, nil
}
