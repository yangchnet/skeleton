-- name: CreateUser :exec
INSERT INTO 
user (username, email, password, created_at)
VALUES ($1, $2, $3, $4);