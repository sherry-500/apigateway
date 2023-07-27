-- name: GetIdlmapByName :one
SELECT * FROM idlmap
WHERE svcname = $1 LIMIT 1;