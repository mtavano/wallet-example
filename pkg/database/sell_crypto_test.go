package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_SellCrypto(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID := "user-id"
	st.createUser(userID)

	inputAmount := int64(1000000) // 10.000

	_, err := st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyBTC,
		Amount:   inputAmount,
	})

	require.NoError(t, err)

	// Act
	newfiatBalance, newCryptoBalance, err := st.SellCrypto(&SellCryptoInput{
		UserID:   userID,
		Currency: CurrencyBTC,
		Amount:   1500,    // 15.00
		Price:    5000000, //50.000
	})

	// Assert
	require.NoError(t, err)
	require.Equal(t, newfiatBalance, int64(75000000))
	require.Equal(t, newCryptoBalance, int64(998500))
}
