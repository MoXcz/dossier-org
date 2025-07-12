-- name: CreateUser :one
INSERT INTO users (name, email, password_hash, role_id)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserFromID :one
SELECT * FROM users
WHERE user_id = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * from users;
