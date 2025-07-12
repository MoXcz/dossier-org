-- name: CreateDossier :one
INSERT INTO dossiers (title, data, assigned_to)
VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetDossiersFromUserID :many
SELECT * FROM dossiers
WHERE assigned_to = $1;

-- name: DeleteDossier :exec
DELETE FROM dossiers;

-- name: GetDossiers :many
SELECT * from dossiers;
