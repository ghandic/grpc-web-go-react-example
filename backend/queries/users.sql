-- name: AddUser :one
INSERT INTO
    users (name)
VALUES
    ( $1 ) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM
    users
WHERE
    id = $1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = $1
LIMIT 1;