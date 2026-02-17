-- name: GetNextFeed :one
SELECT * FROM feeds
ORDER BY last_fetched_at asc NULLS FIRST 
LIMIT 1;