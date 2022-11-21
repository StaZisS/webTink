-- name: CreateUser :one
INSERT INTO users (
    name,
    surname,
    username,
    email,
    grade,
    photo,
    password,
    updated_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8
         )
    RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set name = $2,
    surname = $3,
    username = $4,
    email = $5,
    grade = $6,
    photo = $7,
    verified = $8,
    password = $9,
    role = $10,
    updated_at = $11
WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;