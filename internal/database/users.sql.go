// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const clearFeed = `-- name: ClearFeed :exec
TRUNCATE TABLE feeds
`

func (q *Queries) ClearFeed(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearFeed)
	return err
}

const clearPosts = `-- name: ClearPosts :exec
delete from posts
`

func (q *Queries) ClearPosts(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearPosts)
	return err
}

const clearUser = `-- name: ClearUser :exec
DELETE FROM users
`

func (q *Queries) ClearUser(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, clearUser)
	return err
}

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
	$1,
	$2,
	$3,
	$4,
	$5,
	$6
)
RETURNING id, created_at, updated_at, name, url, user_id, last_fetched_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.Url,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const createFeedFollow = `-- name: CreateFeedFollow :one
with inserted_feed_follow as (
insert into feed_follows (id, created_at, updated_at, user_id, feed_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5
)
returning id, created_at, updated_at, user_id, feed_id
)
select inserted_feed_follow.id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at, inserted_feed_follow.user_id, inserted_feed_follow.feed_id,
feeds.name AS feed_name,
users.name AS user_name
from inserted_feed_follow
inner join feeds
ON inserted_feed_follow.feed_id = feeds.id
inner join users
ON inserted_feed_follow.user_id = users.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const createPost = `-- name: CreatePost :one
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
returning id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, name
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const deleteFollow = `-- name: DeleteFollow :exec
Delete from feed_follows where user_id = $1 and feed_id = $2
`

type DeleteFollowParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) error {
	_, err := q.db.ExecContext(ctx, deleteFollow, arg.UserID, arg.FeedID)
	return err
}

const getFeed = `-- name: GetFeed :one
select id, created_at, updated_at, name, url, user_id, last_fetched_at from feeds where url = $1
`

func (q *Queries) GetFeed(ctx context.Context, url string) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed, url)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :many
select url from feeds
order by last_fetched_at nulls first
limit 5
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getNextFeedToFetch)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		items = append(items, url)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPost = `-- name: GetPost :many
select id, created_at, updated_at, title, url, description, published_at, feed_id from posts
order by published_at
limit $1
`

func (q *Queries) GetPost(ctx context.Context, limit int32) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, getPost, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Title,
			&i.Url,
			&i.Description,
			&i.PublishedAt,
			&i.FeedID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpecPost = `-- name: GetSpecPost :one
Select id, created_at, updated_at, title, url, description, published_at, feed_id from posts
Where url = $1
`

func (q *Queries) GetSpecPost(ctx context.Context, url string) (Post, error) {
	row := q.db.QueryRowContext(ctx, getSpecPost, url)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, created_at, updated_at, name FROM users WHERE name = $1
`

func (q *Queries) GetUser(ctx context.Context, name string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, name)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}

const getUserFollows = `-- name: GetUserFollows :many
select feeds.name from feed_follows
INNER join feeds
on feed_follows.feed_id = feeds.id
where feed_follows.user_id = $1
`

func (q *Queries) GetUserFollows(ctx context.Context, userID uuid.UUID) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFeed = `-- name: ListFeed :many
SELECT feeds.name, url, users.name FROM feeds
INNER JOIN users
ON feeds.user_id = users.id
`

type ListFeedRow struct {
	Name   string
	Url    string
	Name_2 string
}

func (q *Queries) ListFeed(ctx context.Context) ([]ListFeedRow, error) {
	rows, err := q.db.QueryContext(ctx, listFeed)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFeedRow
	for rows.Next() {
		var i ListFeedRow
		if err := rows.Scan(&i.Name, &i.Url, &i.Name_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT name FROM users
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markFeedFetched = `-- name: MarkFeedFetched :exec
update feeds
set updated_at = $1,last_fetched_at = $1
where id = $2
`

type MarkFeedFetchedParams struct {
	UpdatedAt time.Time
	ID        uuid.UUID
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.UpdatedAt, arg.ID)
	return err
}
