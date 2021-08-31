package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_GetLedgerByUserID(t *testing.T) {
	testCases := []struct {
		name             string
		st               *Store
		userID           string
		beforeExpect     func(store *Store, userID string) error
		expectedErrorMsg string
	}{
		{
			name:   "should return correctly without error",
			userID: "user-id",
			st:     NewStore(time.Now),
			beforeExpect: func(store *Store, userID string) error {
				err := store.CreateUser(userID)
				return err
			},
			expectedErrorMsg: "",
		},
		{
			name: "should return an error if user does not exist",
			st:   NewStore(time.Now),
			beforeExpect: func(store *Store, userID string) error {
				return nil
			},
			expectedErrorMsg: ErrUserNotFound.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.beforeExpect(tc.st, tc.userID)

			res, err := tc.st.GetLedgerByUserID(tc.userID)

			if tc.expectedErrorMsg != "" {
				require.Equal(t, tc.expectedErrorMsg, err.Error())
				return
			}

			require.NotNil(t, res)
		})
	}
}
