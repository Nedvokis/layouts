-- name: CreateStaType :one
INSERT INTO sta_types ("bitrix_id", "type")
VALUES ($1, $2)
RETURNING *;