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

-- name: ShowFeeds :many
SELECT feeds.name, feeds.url, users.name AS feed_creator
FROM feeds
INNER JOIN users
ON feeds.user_id = users.id;


-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        $1,
        $2,
        $3,
        (
            SELECT id FROM users
            WHERE users.name = $4
        ),
        (
            SELECT id FROM feeds
            WHERE feeds.url = $5
        )
    )
    RETURNING *
)
SELECT 
    inserted_feed_follow.*, 
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds
    ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users
    ON inserted_feed_follow.user_id = users.id;

-- name: GetFeedByUrl :one
SELECT name FROM feeds
WHERE url = $1;

-- name: GetFeedFollowsByUser :many
SELECT 
    users.name AS user_name, 
    feeds.name AS feed_name 
FROM users
INNER JOIN feed_follows
    ON users.id = feed_follows.user_id
INNER JOIN feeds 
    ON feeds.id = feed_follows.feed_id
WHERE users.name = $1;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows
WHERE feed_follows.user_id = $1 AND feed_follows.feed_id = (
    SELECT id as feed_id FROM feeds
    WHERE url = $2
);
