-- name: CreateRole :one
INSERT INTO roles (name, description)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetRoleFromID :one
SELECT * FROM users
WHERE user_id = $1;

-- name: DeleteRoles :exec
TRUNCATE roles RESTART IDENTITY;

-- name: GetRoles :many
SELECT * from roles;
