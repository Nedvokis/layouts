-- name: CreateStaRoom :one
INSERT INTO sta_room ("bitrix_id", "type")
VALUES ($1, $2)
RETURNING *;