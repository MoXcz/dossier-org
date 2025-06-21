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
DELETE FROM users;

-- name: GetUsers :many
SELECT * from users;
