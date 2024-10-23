-- name: CreateFeedFollow :one

WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)

SELECT 
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedFollowsForUser :many
WITH feed_follow_id AS (
    SELECT feed_id, user_id FROM feed_follows WHERE feed_follows.user_id = $1
)

SELECT 
    feed_follow_id.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follow_id
INNER JOIN feeds ON feed_follow_id.feed_id = feeds.id
INNER JOIN users ON feed_follow_id.user_id = users.id;
