-- name: FeedFollowersForUser :many
Select users.name as user_name, feeds.name as feed_name from feed_follows
INNER JOIN users on user_id = users.id
INNER JOIN feeds on feed_id = feeds.id
WHERE users.name = $1;