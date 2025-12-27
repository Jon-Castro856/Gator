-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, feed_name, url, user_id)
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
SELECT feed_name, url, user_id FROM feeds;

-- name: GetFeedNameAndID :one
SELECT feed_name, id FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET updated_at = NOW(), last_fetched_at = NOW()
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT id, url FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;

