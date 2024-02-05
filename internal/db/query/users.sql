-- name: FindByTGID :one
SELECT *
FROM users
WHERE telegram_id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    user_id, telegram_id, invite_code_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateInviteByTGID :exec
UPDATE users
SET invite_code_id = $2
WHERE telegram_id = $1;

-- name: DeleteByTGID :exec
DELETE FROM users
WHERE telegram_id = $1;