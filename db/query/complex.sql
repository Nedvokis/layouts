-- name: CreateComplex :one
INSERT INTO complexes ("bitrix_id", "name")
VALUES ($1, $2)
RETURNING *;