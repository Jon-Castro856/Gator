-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3
    ) RETURNING *
)
SELECT inserted_feed_follow.*,
    feeds.feed_name AS feed_name,
    users.user_name AS user_name
FROM inserted_feed_follow
INNER JOIN users
ON inserted_feed_follow.user_id = users.id
INNER JOIN feeds
ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT feeds.feed_name, users.id, users.user_name
FROM feed_follows
INNER JOIN feeds
ON feed_follows.feed_id = feeds.id
INNER JOIN users
ON feed_follows.user_id = users.id
WHERE users.id = $1;