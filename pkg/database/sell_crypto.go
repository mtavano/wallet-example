package database

type SellCryptoInput struct {
	UserID   string
	Currency Currency
	Amount   uint64
	Price    uint64
}

func (st *Store) SellCrypto(input *SellCryptoInput) (int64, int64, error) {
	newCryptoBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Currency: input.Currency,
		Amount:   -1 * int64(input.Amount), // we are loading cents again
	})
	if err != nil {
		return 0, 0, err
	}

	fiatToLoad := float64(input.Amount/100) * float64(input.Price/100)
	newFiatBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Currency: CurrencyUSDC,
		Amount:   int64(fiatToLoad * 100), // we are loading cents again
	})
	if err != nil {
		return 0, 0, err
	}

	return newFiatBalance, newCryptoBalance, nil
}
