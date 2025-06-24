-- name: CreateUser :one
INSERT INTO users (name, email, encryptedPassword)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserFromID :one
SELECT * FROM users
WHERE id = $1;

-- name: DeleteUsers :exec
TRUNCATE users RESTART IDENTITY;

-- name: GetUsers :many
SELECT * from users;
