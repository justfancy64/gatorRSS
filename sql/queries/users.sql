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
DELETE FROM users;
-- name: ClearFeed :exec
TRUNCATE TABLE feeds;

-- name: ListUsers :many
SELECT name FROM users
ORDER BY name;

-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING *;



-- name: ListFeed :many
SELECT feeds.name, url, users.name FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;
