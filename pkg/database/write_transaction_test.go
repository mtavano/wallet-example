package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_WriteTransaction(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID := "user-id"
	st.CreateUser(userID)

	inputAmount := int64(100)

	// Act
	newBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyUSDC,
		Amount:   inputAmount,
	})

	require.NoError(t, err)

	// Assert
	require.Equal(t, newBalance, inputAmount)
}

func Test_WriteTransaction_NewBalance(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID := "user-id"
	st.CreateUser(userID)

	inputAmount := int64(100)

	// Act
	newBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyUSDC,
		Amount:   3 * inputAmount,
	})

	require.Equal(t, newBalance, 3*inputAmount)

	require.NoError(t, err)

	newBalance, err = st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyUSDC,
		Amount:   -1 * inputAmount,
	})

	// Assert
	require.Equal(t, newBalance, 2*inputAmount)
}

func Test_WriteTransaction_InsuficientFunds_Error(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID := "user-id"
	st.CreateUser(userID)

	inputAmount := int64(-100)

	// Act
	newBalance, err := st.writeTransaction(&writeTransactionInput{
		UserID:   userID,
		Currency: CurrencyUSDC,
		Amount:   inputAmount,
	})

	require.Equal(t, ErrInsuficientFunds.Error(), err.Error())

	// Assert
	require.Equal(t, newBalance, int64(0))
}
