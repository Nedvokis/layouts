-- name: CreateLitter :one
INSERT INTO litters ("parent", "bitrix_id", "name")
VALUES ($1, $2, $3)
RETURNING *;