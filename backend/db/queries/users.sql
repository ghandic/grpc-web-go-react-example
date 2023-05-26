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

-- name: GetUsers :many
SELECT *
FROM users
WHERE 
    (CASE WHEN @search <> '' THEN name ILIKE CONCAT('%', @search, '%') ELSE TRUE END)
ORDER BY 
    CASE WHEN @created_at_asc::bool THEN created_at END asc,
    CASE WHEN @created_at_desc::bool THEN created_at END desc,
    CASE WHEN @id_asc::bool THEN id END asc,
    CASE WHEN @id_desc::bool THEN id END desc,
    CASE WHEN @name_asc::bool THEN name END asc,
    CASE WHEN @name_desc::bool THEN name END desc
OFFSET NULLIF(@offset_amount::int, 0)
LIMIT NULLIF(@limit_amount::int, 0);

-- name: GetUsersCount :one
SELECT COUNT(*)
FROM users
WHERE 
    (CASE WHEN @search <> '' THEN name ILIKE CONCAT('%', @search, '%') ELSE TRUE END);

-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = $1
LIMIT 1;