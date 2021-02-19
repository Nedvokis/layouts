-- name: CreateStaType :one
INSERT INTO sta_types ("bitrix_id", "type_name")
VALUES ($1, $2)
RETURNING *;