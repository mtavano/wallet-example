package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_NewStore(t *testing.T) {
	st := NewStore(time.Now)

	require.NotNil(t, st)
}
