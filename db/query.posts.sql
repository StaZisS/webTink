-- name: CreatePost :one
INSERT INTO posts (
    title,
    content,
    author_id,
    photo,
    updated_at
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING *;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
set title = $2,
    content = $3,
    photo = $4,
WHERE id = $1
    RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;