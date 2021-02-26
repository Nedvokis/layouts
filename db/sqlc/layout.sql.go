// Code generated by sqlc. DO NOT EDIT.
// source: layout.sql

package db

import (
	"context"
	"database/sql"
)

const createLayout = `-- name: CreateLayout :one
INSERT INTO layouts (
		"parent",
		"area",
		"citchen_area",
		"door",
		"floor",
		"bitrix_id",
		"layout_id",
		"living_area",
		"num",
		"price",
		"room",
		"status",
		"layouts_url",
		"svg_path",
		"type"
	)
VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15
	)
RETURNING id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
`

type CreateLayoutParams struct {
	Parent      int64           `json:"parent"`
	Area        sql.NullFloat64 `json:"area"`
	CitchenArea sql.NullFloat64 `json:"citchen_area"`
	Door        sql.NullInt32   `json:"door"`
	Floor       sql.NullInt32   `json:"floor"`
	BitrixID    sql.NullInt32   `json:"bitrix_id"`
	LayoutID    sql.NullInt32   `json:"layout_id"`
	LivingArea  sql.NullFloat64 `json:"living_area"`
	Num         sql.NullString  `json:"num"`
	Price       sql.NullInt32   `json:"price"`
	Room        sql.NullInt32   `json:"room"`
	Status      sql.NullInt32   `json:"status"`
	LayoutsUrl  sql.NullString  `json:"layouts_url"`
	SvgPath     sql.NullString  `json:"svg_path"`
	Type        sql.NullInt32   `json:"type"`
}

func (q *Queries) CreateLayout(ctx context.Context, arg CreateLayoutParams) (Layout, error) {
	row := q.db.QueryRowContext(ctx, createLayout,
		arg.Parent,
		arg.Area,
		arg.CitchenArea,
		arg.Door,
		arg.Floor,
		arg.BitrixID,
		arg.LayoutID,
		arg.LivingArea,
		arg.Num,
		arg.Price,
		arg.Room,
		arg.Status,
		arg.LayoutsUrl,
		arg.SvgPath,
		arg.Type,
	)
	var i Layout
	err := row.Scan(
		&i.ID,
		&i.Parent,
		&i.Area,
		&i.CitchenArea,
		&i.Door,
		&i.Floor,
		&i.BitrixID,
		&i.LayoutID,
		&i.LivingArea,
		&i.Num,
		&i.Price,
		&i.Status,
		&i.Type,
		&i.Room,
		&i.LayoutsUrl,
		&i.SvgPath,
	)
	return i, err
}

const getAllListLayouts = `-- name: GetAllListLayouts :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
`

func (q *Queries) GetAllListLayouts(ctx context.Context) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getAllListLayouts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Layout
	for rows.Next() {
		var i Layout
		if err := rows.Scan(
			&i.ID,
			&i.Parent,
			&i.Area,
			&i.CitchenArea,
			&i.Door,
			&i.Floor,
			&i.BitrixID,
			&i.LayoutID,
			&i.LivingArea,
			&i.Num,
			&i.Price,
			&i.Status,
			&i.Type,
			&i.Room,
			&i.LayoutsUrl,
			&i.SvgPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLayout = `-- name: GetLayout :one
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetLayout(ctx context.Context, id int64) (Layout, error) {
	row := q.db.QueryRowContext(ctx, getLayout, id)
	var i Layout
	err := row.Scan(
		&i.ID,
		&i.Parent,
		&i.Area,
		&i.CitchenArea,
		&i.Door,
		&i.Floor,
		&i.BitrixID,
		&i.LayoutID,
		&i.LivingArea,
		&i.Num,
		&i.Price,
		&i.Status,
		&i.Type,
		&i.Room,
		&i.LayoutsUrl,
		&i.SvgPath,
	)
	return i, err
}

const getLayoutByLitter = `-- name: GetLayoutByLitter :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
WHERE parent = $1
	AND door = $2
`

type GetLayoutByLitterParams struct {
	Parent int64         `json:"parent"`
	Door   sql.NullInt32 `json:"door"`
}

func (q *Queries) GetLayoutByLitter(ctx context.Context, arg GetLayoutByLitterParams) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getLayoutByLitter, arg.Parent, arg.Door)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Layout
	for rows.Next() {
		var i Layout
		if err := rows.Scan(
			&i.ID,
			&i.Parent,
			&i.Area,
			&i.CitchenArea,
			&i.Door,
			&i.Floor,
			&i.BitrixID,
			&i.LayoutID,
			&i.LivingArea,
			&i.Num,
			&i.Price,
			&i.Status,
			&i.Type,
			&i.Room,
			&i.LayoutsUrl,
			&i.SvgPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getListLayouts = `-- name: GetListLayouts :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
ORDER BY name
LIMIT $1 OFFSET $2
`

type GetListLayoutsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListLayouts(ctx context.Context, arg GetListLayoutsParams) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getListLayouts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Layout
	for rows.Next() {
		var i Layout
		if err := rows.Scan(
			&i.ID,
			&i.Parent,
			&i.Area,
			&i.CitchenArea,
			&i.Door,
			&i.Floor,
			&i.BitrixID,
			&i.LayoutID,
			&i.LivingArea,
			&i.Num,
			&i.Price,
			&i.Status,
			&i.Type,
			&i.Room,
			&i.LayoutsUrl,
			&i.SvgPath,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSvgPath = `-- name: UpdateSvgPath :exec
UPDATE layouts
SET svg_path = $2
WHERE id = $1
`

type UpdateSvgPathParams struct {
	ID      int64          `json:"id"`
	SvgPath sql.NullString `json:"svg_path"`
}

func (q *Queries) UpdateSvgPath(ctx context.Context, arg UpdateSvgPathParams) error {
	_, err := q.db.ExecContext(ctx, updateSvgPath, arg.ID, arg.SvgPath)
	return err
}
