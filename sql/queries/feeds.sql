-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updatable_at, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFeeds :many
SELECT *
FROM feeds;