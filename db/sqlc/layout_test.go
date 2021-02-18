package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLayout(t *testing.T) {
	arg := CreateLayoutParams{
		Parent:      1,
		Area:        sql.NullFloat64{Float64: 5.69, Valid: true},
		CitchenArea: sql.NullFloat64{Float64: 1.69, Valid: true},
		Door:        sql.NullInt32{Int32: 420, Valid: true},
		Floor:       sql.NullInt32{Int32: 6, Valid: true},
		BitrixID:    sql.NullInt32{Int32: 1, Valid: true},
		LayoutID:    sql.NullInt32{Int32: 1, Valid: true},
		LivingArea:  sql.NullFloat64{Float64: 3.69, Valid: true},
		Num:         sql.NullInt32{Int32: 1, Valid: true},
		Price:       sql.NullInt32{Int32: 3500, Valid: true},
		Room:        sql.NullInt32{Int32: 2, Valid: true},
		Status:      sql.NullInt32{Int32: 4, Valid: true},
		SvgPath:     sql.NullString{},
		Type:        sql.NullInt32{Int32: 3, Valid: true},
	}

	layout, err := testQueries.CreateLayout(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, layout)

	require.Equal(t, arg.Area, layout.Area)
}
