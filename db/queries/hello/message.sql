-- name: CreateMessage :exec
INSERT INTO message (   
    id, 
    sender, 
    message, 
    created_at
)VALUES (
    $1, $2, $3, $4
);