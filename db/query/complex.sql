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