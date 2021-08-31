package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_CreateTransfer(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID1 := "user-id-1"
	userID2 := "user-id-2"

	err := st.createUser(userID1)
	require.NoError(t, err)

	err = st.createUser(userID2)
	require.NoError(t, err)

	inputAmount := int64(100000000) // 1.000.000

	_, err = st.writeTransaction(&writeTransactionInput{
		UserID:   userID1,
		Currency: CurrencyUSDC,
		Amount:   inputAmount,
	})

	require.NoError(t, err)

	// Act

	newSrcBalance, newDestBalance, err := st.CreateTransfer(&CreateTransferInput{
		SourceID:      userID1,
		DestinationID: userID2,
		Currency:      CurrencyUSDC,
		Amount:        100000000, // 1.000.000
	})

	require.NoError(t, err)

	require.Equal(t, newSrcBalance, int64(0))
	require.Equal(t, newDestBalance, int64(inputAmount))
}

func Test_CreateTransfer_ErrUserNotFound(t *testing.T) {
	// Arrange
	now := time.Now()
	st := NewStore(func() time.Time { return now })

	userID1 := "user-id-1"
	userID2 := "user-id-2"

	err := st.createUser(userID1)
	require.NoError(t, err)
	inputAmount := int64(100000000) // 1.000.000

	_, err = st.writeTransaction(&writeTransactionInput{
		UserID:   userID1,
		Currency: CurrencyUSDC,
		Amount:   inputAmount,
	})

	require.NoError(t, err)

	// Act

	_, _, err = st.CreateTransfer(&CreateTransferInput{
		SourceID:      userID1,
		DestinationID: userID2,
		Currency:      CurrencyUSDC,
		Amount:        100000000, // 1.000.000
	})

	require.Equal(t, err.Error(), ErrUserNotFound.Error())
}
