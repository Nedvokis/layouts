-- name: CreateLayout :one
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
RETURNING *;
-- name: GetLayout :one
SELECT *
FROM layouts
WHERE id = $1
LIMIT 1;
-- name: GetLayoutByLitter :many
SELECT *
FROM layouts
WHERE parent = $1
	AND type = 1
	AND status = 2;
-- name: GetLayoutByLitterAndDoor :many
SELECT *
FROM layouts
WHERE parent = $1
	AND door = $2
	AND type = 1
	AND status = 2;
-- name: GetListLayouts :many
SELECT *
FROM layouts
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: GetAllListLayouts :many
SELECT *
FROM layouts;
-- name: UpdateSvgPath :exec
UPDATE layouts
SET svg_path = $2
WHERE id = $1;