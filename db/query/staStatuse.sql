-- name: CreateStaStatuse :one
INSERT INTO sta_statuses ("bitrix_id", "type_name")
VALUES ($1, $2)
RETURNING *;