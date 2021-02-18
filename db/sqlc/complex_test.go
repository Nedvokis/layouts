package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateComplex(t *testing.T) {
	arg := CreateComplexParams{
		BitrixID: 2,
		Name:     sql.NullString{String: "ЖК \"Движение\"", Valid: true},
	}

	complex, err := testQueries.CreateComplex(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, complex)

	require.Equal(t, arg.Name, complex.Name)
}
