-- name: CreateComplex :one
INSERT INTO complexes ("bitrix_id", "name")
VALUES ($1, $2)
RETURNING *;
-- name: GetComplex :one
SELECT *
FROM complexes
WHERE id = $1
LIMIT 1;
-- name: GetListComplex :many
SELECT *
FROM complexes
ORDER BY name
LIMIT $1 OFFSET $2;
-- name: GetListAllComplexes :many
SELECT *
FROM complexes;
-- name: GetComplexByBxID :one
SELECT *
FROM complexes
WHERE bitrix_id = @bitrix_id
LIMIT 1;