-- name: FeedByUrl :one
SELECT id,name FROM feeds 
WHERE url = $1;