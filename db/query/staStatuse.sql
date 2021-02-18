-- name: CreateStaStatuse :one
INSERT INTO sta_statuses ("bitrix_id", "type")
VALUES ($1, $2)
RETURNING *;