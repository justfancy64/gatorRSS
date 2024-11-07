-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUser :one
select * FROM users WHERE name = $1;


-- name: ClearUser :exec
TRUNCATE TABLE users;

-- name: ListUsers :many
SELECT name FROM users
ORDER BY name;
