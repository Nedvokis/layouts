// Code generated by sqlc. DO NOT EDIT.
// source: staStatuse.sql

package db

import (
	"context"
	"database/sql"
)

const createStaStatuse = `-- name: CreateStaStatuse :one
INSERT INTO sta_statuses ("bitrix_id", "type")
VALUES ($1, $2)
RETURNING id, bitrix_id, type
`

type CreateStaStatuseParams struct {
	BitrixID int64          `json:"bitrix_id"`
	Type     sql.NullString `json:"type"`
}

func (q *Queries) CreateStaStatuse(ctx context.Context, arg CreateStaStatuseParams) (StaStatus, error) {
	row := q.db.QueryRowContext(ctx, createStaStatuse, arg.BitrixID, arg.Type)
	var i StaStatus
	err := row.Scan(&i.ID, &i.BitrixID, &i.Type)
	return i, err
}
