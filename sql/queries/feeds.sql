-- name: CreateFeed :one
INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id
)
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
SELECT name, url, user_id FROM feeds;

-- name: GetFeedIDFromURL :one
SELECT id FROM feeds WHERE url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = $1, updated_at = $2
WHERE id = $3 AND user_id = $4;

-- name: GetNextFeedToFetch :one
SELECT id, url
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;
