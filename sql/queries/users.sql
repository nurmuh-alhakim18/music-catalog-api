-- name: CreateUser :one
INSERT INTO users (id, email, username, password)
VALUES (
  gen_random_uuid(), $1, $2, $3
)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;