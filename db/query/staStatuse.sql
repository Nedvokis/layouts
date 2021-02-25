-- name: CreateStaStatuse :one
INSERT INTO sta_statuses ("bitrix_id", "type_name")
VALUES ($1, $2)
RETURNING *;
-- name: GetListAllStaStatuse :many
SELECT *
FROM sta_statuses
ORDER BY id;