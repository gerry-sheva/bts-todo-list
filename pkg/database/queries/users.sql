-- name: NewUser :one
insert into users (username, password)
values ($1, $2)
returning username;

-- name: GetUser :one
SELECT username, password FROM users
WHERE username = $1;
