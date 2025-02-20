-- name: NewUser :one
insert into users (username, password)
values ($1, $2)
returning user_id;

-- name: GetUser :one
SELECT user_id, password FROM users
WHERE username = $1;
