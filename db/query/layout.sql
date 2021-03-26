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
	AND status = 2
LIMIT 12;
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
-- name: GetFilteredLayouts :many
SELECT *
FROM layouts
WHERE type = 1
	AND status = 2
	AND (
		CASE
			WHEN room = @room::int
			OR 0 = @room::int THEN true
		END
	)
	AND (
		CASE
			WHEN parent = ANY(@parent::int [])
			OR 0 = ANY(@parent::int [0]) THEN true
		END
	)
	AND (
		CASE
			WHEN area >= @area_min::float
			AND area <= @area_max::float THEN true
		END
	)
	AND (
		CASE
			WHEN living_area >= @living_area_min::float
			AND living_area <= @living_area_max::float THEN true
		END
	)
	AND (
		CASE
			WHEN citchen_area >= @citchen_area_min::float
			AND citchen_area <= @citchen_area_max::float THEN true
		END
	)
ORDER BY (
		CASE
			WHEN @citchen_area_desc::bool THEN citchen_area
		END
	) desc,
	(
		CASE
			WHEN @citchen_area_asc::bool THEN citchen_area
		END
	) asc,
	(
		CASE
			WHEN @living_area_desc::bool THEN living_area
		END
	) desc,
	(
		CASE
			WHEN @living_area_asc::bool THEN living_area
		END
	) asc,
	(
		CASE
			WHEN @area_desc::bool THEN area
		END
	) desc,
	(
		CASE
			WHEN @area_asc::bool THEN area
		END
	) asc OFFSET @off_set::int
LIMIT 12;
-- name: GetFilteredLayoutsLength :many
SELECT *
FROM layouts
WHERE type = 1
	AND status = 2
	AND (
		CASE
			WHEN room = @room::int
			OR 0 = @room::int THEN true
		END
	)
	AND (
		CASE
			WHEN parent = ANY(@parent::int [])
			OR 0 = ANY(@parent::int []) THEN true
		END
	)
	AND (
		CASE
			WHEN area >= @area_min::float
			AND area <= @area_max::float THEN true
		END
	)
	AND (
		CASE
			WHEN living_area >= @living_area_min::float
			AND living_area <= @living_area_max::float THEN true
		END
	)
	AND (
		CASE
			WHEN citchen_area >= @citchen_area_min::float
			AND citchen_area <= @citchen_area_max::float THEN true
		END
	)
ORDER BY (
		CASE
			WHEN @citchen_area_desc::bool THEN citchen_area
		END
	) desc,
	(
		CASE
			WHEN @citchen_area_asc::bool THEN citchen_area
		END
	) asc,
	(
		CASE
			WHEN @living_area_desc::bool THEN living_area
		END
	) desc,
	(
		CASE
			WHEN @living_area_asc::bool THEN living_area
		END
	) asc,
	(
		CASE
			WHEN @area_desc::bool THEN area
		END
	) desc,
	(
		CASE
			WHEN @area_asc::bool THEN area
		END
	) asc;