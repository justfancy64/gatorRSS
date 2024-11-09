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
SELECT * FROM users WHERE name = $1;


-- name: ClearUser :exec
DELETE FROM users;
-- name: ClearFeed :exec
TRUNCATE TABLE feeds;

-- name: ClearPosts :exec
delete from posts;

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



-- name: CreateFeedFollow :one
with inserted_feed_follow as (
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5
)
returning *
)
select inserted_feed_follow.*,
feeds.name AS feed_name,
users.name AS user_name
from inserted_feed_follow
inner join feeds
ON inserted_feed_follow.feed_id = feeds.id
inner join users
ON inserted_feed_follow.user_id = users.id;


-- name: GetFeed :one
select * from feeds where url = $1;


-- name: GetUserFollows :many
select feeds.name from feed_follows
INNER join feeds
on feed_follows.feed_id = feeds.id
where feed_follows.user_id = $1;

-- name: DeleteFollow :exec
Delete from feed_follows where user_id = $1 and feed_id = $2;


-- name: MarkFeedFetched :exec
update feeds
set updated_at = $1,last_fetched_at = $1
where id = $2;

-- name: GetNextFeedToFetch :one
select url from feeds
order by last_fetched_at nulls first
limit 1;



-- name: CreatePost :one
insert into posts (id,created_at,updated_at,title,url,description,published_at,feed_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
    )
returning *;


-- name: GetPost :many
select * from posts
order by published_at
limit $1;




