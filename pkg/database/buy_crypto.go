package database

import (
	"github.com/pkg/errors"
)

type BuyCryptoInput struct {
	UserID   string
	Currency Currency
	Amount   uint64
	Price    uint64
}

func (st *Store) BuyCrypto(input *BuyCryptoInput) (int64, int64, error) {
	// amount * price
	// we are dividing by 100 to avoid cents issue
	fiatAmountToSpent := float64(input.Amount/100) * float64(input.Price/100)
	newFiatBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Currency: CurrencyUSDC,
		Amount:   int64(-1 * fiatAmountToSpent * 100), // we are loading cents again
	})
	if err != nil {
		return 0, 0, errors.Wrap(err, "wrapped: error")
	}

	newCryptoBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   input.UserID,
		Currency: input.Currency,
		Amount:   int64(input.Amount),
	})
	if err != nil {
		return 0, 0, err
	}

	return newFiatBalance, newCryptoBalance, nil
}
