package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BuyCrypto(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID := "user-id"
	st.createUser(userID)

	inputAmount := int64(100000000) // 1.000.000

	_, err := st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyUSDC,
		Amount:   inputAmount,
	})

	require.NoError(t, err)

	// Act

	newfiatBalance, newCryptoBalance, err := st.BuyCrypto(&BuyCryptoInput{
		UserID:   userID,
		Currency: CurrencyBTC,
		Amount:   1500,    // 15.00
		Price:    5000000, //50.000
	})

	// Assert
	require.NoError(t, err)
	require.Equal(t, newfiatBalance, int64(25000000))
	require.Equal(t, newCryptoBalance, int64(1500))
}
