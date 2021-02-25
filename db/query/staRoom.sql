-- name: CreateStaRoom :one
INSERT INTO sta_room ("bitrix_id", "type_name")
VALUES ($1, $2)
RETURNING *;
-- name: GetListAllStaRoom :many
SELECT *
FROM sta_room
ORDER BY id;