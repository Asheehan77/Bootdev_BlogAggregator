-- name: GetPostsByUser :many
SELECT posts.* FROM posts
INNER JOIN feed_follows on posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY published_at DESC
LIMIT $2;
