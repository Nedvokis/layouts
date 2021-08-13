// Code generated by sqlc. DO NOT EDIT.
// source: layout.sql

package db

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
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

const getFilteredLayouts = `-- name: GetFilteredLayouts :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
WHERE type = 1
	AND status = 2
	AND (
		CASE
			WHEN bitrix_id = $1::int
			OR 0 = ANY($1::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN room = ANY($2::int [])
			OR 0 = ANY($2::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN parent = ANY($3::int [])
			OR 0 = ANY($3::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN area >= $4::float
			AND area <= $5::float THEN true
		END
	)
	AND (
		CASE
			WHEN living_area >= $6::float
			AND living_area <= $7::float THEN true
		END
	)
	AND (
		CASE
			WHEN citchen_area >= $8::float
			AND citchen_area <= $9::float THEN true
		END
	)
ORDER BY (
		CASE
			WHEN $10::bool THEN citchen_area
		END
	) desc,
	(
		CASE
			WHEN $11::bool THEN citchen_area
		END
	) asc,
	(
		CASE
			WHEN $12::bool THEN living_area
		END
	) desc,
	(
		CASE
			WHEN $13::bool THEN living_area
		END
	) asc,
	(
		CASE
			WHEN $14::bool THEN area
		END
	) desc,
	(
		CASE
			WHEN $15::bool THEN area
		END
	) asc OFFSET $16::float
LIMIT 12
`

type GetFilteredLayoutsParams struct {
	BitrixID        int32   `json:"bitrix_id"`
	Room            []int32 `json:"room"`
	Parent          []int32 `json:"parent"`
	AreaMin         float64 `json:"area_min"`
	AreaMax         float64 `json:"area_max"`
	LivingAreaMin   float64 `json:"living_area_min"`
	LivingAreaMax   float64 `json:"living_area_max"`
	CitchenAreaMin  float64 `json:"citchen_area_min"`
	CitchenAreaMax  float64 `json:"citchen_area_max"`
	CitchenAreaDesc bool    `json:"citchen_area_desc"`
	CitchenAreaAsc  bool    `json:"citchen_area_asc"`
	LivingAreaDesc  bool    `json:"living_area_desc"`
	LivingAreaAsc   bool    `json:"living_area_asc"`
	AreaDesc        bool    `json:"area_desc"`
	AreaAsc         bool    `json:"area_asc"`
	OffSet          float64 `json:"off_set"`
}

func (q *Queries) GetFilteredLayouts(ctx context.Context, arg GetFilteredLayoutsParams) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getFilteredLayouts,
		arg.BitrixID,
		pq.Array(arg.Room),
		pq.Array(arg.Parent),
		arg.AreaMin,
		arg.AreaMax,
		arg.LivingAreaMin,
		arg.LivingAreaMax,
		arg.CitchenAreaMin,
		arg.CitchenAreaMax,
		arg.CitchenAreaDesc,
		arg.CitchenAreaAsc,
		arg.LivingAreaDesc,
		arg.LivingAreaAsc,
		arg.AreaDesc,
		arg.AreaAsc,
		arg.OffSet,
	)
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

const getFilteredLayoutsLength = `-- name: GetFilteredLayoutsLength :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
WHERE type = 1
	AND status = 2
	AND (
		CASE
			WHEN room = ANY($1::int [])
			OR 0 = ANY($1::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN bitrix_id = $2::int
			OR 0 = ANY($2::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN parent = ANY($3::int [])
			OR 0 = ANY($3::int [1]) THEN true
		END
	)
	AND (
		CASE
			WHEN area >= $4::float
			AND area <= $5::float THEN true
		END
	)
	AND (
		CASE
			WHEN living_area >= $6::float
			AND living_area <= $7::float THEN true
		END
	)
	AND (
		CASE
			WHEN citchen_area >= $8::float
			AND citchen_area <= $9::float THEN true
		END
	)
ORDER BY (
		CASE
			WHEN $10::bool THEN citchen_area
		END
	) desc,
	(
		CASE
			WHEN $11::bool THEN citchen_area
		END
	) asc,
	(
		CASE
			WHEN $12::bool THEN living_area
		END
	) desc,
	(
		CASE
			WHEN $13::bool THEN living_area
		END
	) asc,
	(
		CASE
			WHEN $14::bool THEN area
		END
	) desc,
	(
		CASE
			WHEN $15::bool THEN area
		END
	) asc
`

type GetFilteredLayoutsLengthParams struct {
	Room            []int32 `json:"room"`
	BitrixID        int32   `json:"bitrix_id"`
	Parent          []int32 `json:"parent"`
	AreaMin         float64 `json:"area_min"`
	AreaMax         float64 `json:"area_max"`
	LivingAreaMin   float64 `json:"living_area_min"`
	LivingAreaMax   float64 `json:"living_area_max"`
	CitchenAreaMin  float64 `json:"citchen_area_min"`
	CitchenAreaMax  float64 `json:"citchen_area_max"`
	CitchenAreaDesc bool    `json:"citchen_area_desc"`
	CitchenAreaAsc  bool    `json:"citchen_area_asc"`
	LivingAreaDesc  bool    `json:"living_area_desc"`
	LivingAreaAsc   bool    `json:"living_area_asc"`
	AreaDesc        bool    `json:"area_desc"`
	AreaAsc         bool    `json:"area_asc"`
}

func (q *Queries) GetFilteredLayoutsLength(ctx context.Context, arg GetFilteredLayoutsLengthParams) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getFilteredLayoutsLength,
		pq.Array(arg.Room),
		arg.BitrixID,
		pq.Array(arg.Parent),
		arg.AreaMin,
		arg.AreaMax,
		arg.LivingAreaMin,
		arg.LivingAreaMax,
		arg.CitchenAreaMin,
		arg.CitchenAreaMax,
		arg.CitchenAreaDesc,
		arg.CitchenAreaAsc,
		arg.LivingAreaDesc,
		arg.LivingAreaAsc,
		arg.AreaDesc,
		arg.AreaAsc,
	)
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
LIMIT 12
`

func (q *Queries) GetLayoutByLitter(ctx context.Context, parent int64) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getLayoutByLitter, parent)
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

const getLayoutByLitterAndDoor = `-- name: GetLayoutByLitterAndDoor :many
SELECT id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
FROM layouts
WHERE parent = $1
	AND door = $2
`

type GetLayoutByLitterAndDoorParams struct {
	Parent int64         `json:"parent"`
	Door   sql.NullInt32 `json:"door"`
}

func (q *Queries) GetLayoutByLitterAndDoor(ctx context.Context, arg GetLayoutByLitterAndDoorParams) ([]Layout, error) {
	rows, err := q.db.QueryContext(ctx, getLayoutByLitterAndDoor, arg.Parent, arg.Door)
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

const updateLayout = `-- name: UpdateLayout :one
UPDATE layouts
SET area = $1,
	citchen_area = $2,
	door = $3,
	floor = $4,
	living_area = $5,
	num = $6,
	price = $7,
	room = $8,
	status = $9,
	layouts_url = $10,
	type = $11
WHERE bitrix_id = $12
RETURNING id, parent, area, citchen_area, door, floor, bitrix_id, layout_id, living_area, num, price, status, type, room, layouts_url, svg_path
`

type UpdateLayoutParams struct {
	Area        sql.NullFloat64 `json:"area"`
	CitchenArea sql.NullFloat64 `json:"citchen_area"`
	Door        sql.NullInt32   `json:"door"`
	Floor       sql.NullInt32   `json:"floor"`
	LivingArea  sql.NullFloat64 `json:"living_area"`
	Num         sql.NullString  `json:"num"`
	Price       sql.NullInt32   `json:"price"`
	Room        sql.NullInt32   `json:"room"`
	Status      sql.NullInt32   `json:"status"`
	LayoutsUrl  sql.NullString  `json:"layouts_url"`
	Type        sql.NullInt32   `json:"type"`
	BitrixID    sql.NullInt32   `json:"bitrix_id"`
}

func (q *Queries) UpdateLayout(ctx context.Context, arg UpdateLayoutParams) (Layout, error) {
	row := q.db.QueryRowContext(ctx, updateLayout,
		arg.Area,
		arg.CitchenArea,
		arg.Door,
		arg.Floor,
		arg.LivingArea,
		arg.Num,
		arg.Price,
		arg.Room,
		arg.Status,
		arg.LayoutsUrl,
		arg.Type,
		arg.BitrixID,
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
