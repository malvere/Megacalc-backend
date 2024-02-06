-- name: FindCode :one
SELECT *
FROM invite_codes
WHERE code = $1 LIMIT 1;

-- name: ListAllCodes :many
SELECT *
FROM invite_codes
LIMIT 20
OFFSET $1;

-- name: CreateCode :one
INSERT INTO invite_codes (
    code_id, code, active
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateCode :exec
UPDATE invite_codes
SET active = $2
WHERE code = $1;

-- name: DeleteCode :exec
DELETE FROM invite_codes
WHERE code = $1;