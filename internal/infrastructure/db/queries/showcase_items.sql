-- name: GetShowcaseItem :one
SELECT * FROM showcase_item WHERE id = $1 LIMIT 1;

-- name: ListShowcaseItem :many
SELECT * FROM showcase_item ORDER BY title;
