-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetUserFromFeed :one
SELECT users.name
FROM users
INNER JOIN feeds
ON users.id = feeds.user_id
WHERE feeds.id = $1;