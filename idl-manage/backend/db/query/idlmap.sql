-- name: CreateIdlmap :one
INSERT INTO idlmap (
    svcname,
    idl,
    type
) VALUES (
    $1, $2, $3
)RETURNING *;

-- name: GetIdlmap :one
SELECT * FROM idlmap
WHERE id = $1 LIMIT 1;

-- name: GetIdlmapByName :one
SELECT * FROM idlmap
WHERE svcname = $1 LIMIT 1;

-- name: ListIdlmap :many
SELECT * FROM idlmap
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateIdlmap :one
UPDATE idlmap
SET idl = $2, type = $3
WHERE svcname = $1
RETURNING *;

-- name: DeleteIdlmap :exec
DELETE FROM idlmap
WHERE svcname = $1;