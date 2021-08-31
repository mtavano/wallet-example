package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_Store_CreateUser(t *testing.T) {
	st := NewStore(time.Now)
	require.Equal(t, len(st.userLedgers), 0)
	userID := "user-id-001"

	err := st.CreateUser(userID)

	require.NoError(t, err)

	require.Equal(t, len(st.userLedgers), 1)

	ledgerCurrency := st.userLedgers[UserID(userID)]

	require.Equal(t, len(ledgerCurrency.BalancesCurrency), len(supportedCurrencies))
	require.Equal(t, len(ledgerCurrency.TransactionsCurrency), len(supportedCurrencies))

	for currency, _ := range supportedCurrencies {
		txs := ledgerCurrency.TransactionsCurrency[currency]
		require.Equal(t, len(txs), 0)
		require.Equal(t, ledgerCurrency.BalancesCurrency[currency], int64(0))
	}
}

func Test_Store_CreateUser_Error_UserAlreadyExist(t *testing.T) {
	expectedErrMsg := "database: Store.CreateUser user already exist"

	st := NewStore(time.Now)
	userID := "user-id-001"

	err := st.CreateUser(userID)

	require.Nil(t, err)

	err = st.CreateUser(userID)

	require.Equal(t, expectedErrMsg, err.Error())
}
