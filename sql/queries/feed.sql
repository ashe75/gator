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


-- name: GetUserNameCreatedFeed :many
SELECT feeds.name, feeds.url, users.name 
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedsByURL :one
SELECT * FROM feeds
WHERE url = $1;

-- name: MarkFeedFetched :one
UPDATE feeds
SET updated_at = $1, last_fetched_at = $2
WHERE id = $3
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;