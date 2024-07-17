-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetFeeds :many
SELECT *
FROM feeds;
-- name: GetLastFeeds :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;
-- name: MarkFeedFetched :many
UPDATE feeds
SET updated_at = NOW(),
    last_fetched_at = NOW()
WHERE id = $1
RETURNING *;