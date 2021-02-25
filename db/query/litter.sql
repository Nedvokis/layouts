-- name: CreateLitter :one
INSERT INTO litters ("parent", "bitrix_id", "name")
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetLitters :one
SELECT *
FROM litters
WHERE id = $1
LIMIT 1;
-- name: GetListLitters :many
SELECT *
FROM litters
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: GetListAllLitters :many
SELECT *
FROM litters
ORDER BY name;