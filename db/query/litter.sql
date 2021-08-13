-- name: CreateLitter :one
INSERT INTO litters ("parent", "bitrix_id", "name")
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetLitter :one
SELECT *
FROM litters
WHERE id = $1
LIMIT 1;
-- name: GetListLitters :many
SELECT *
FROM litters
ORDER BY id;
-- name: GetListAllLitters :many
SELECT *
FROM litters;
-- name: GetListLittersByParent :many
SELECT *
FROM litters
WHERE parent = $1
ORDER BY id;
-- name: GetLitterByBxID :one
SELECT *
FROM litters
WHERE bitrix_id = @bitrix_id
LIMIT 1;