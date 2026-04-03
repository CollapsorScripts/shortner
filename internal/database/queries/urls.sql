-- name: CreateUrl :one
insert into urls (original_url, short_url) values ($1, $2) ON CONFLICT (short_url) DO NOTHING returning *;

-- name: GetOriginalUrlByShortUrl :one
select original_url from urls where short_url = $1;

-- name: GetOriginalUrlById :one
select original_url from urls where id = $1;

-- name: GetUrlByShortUrl :one
select * from urls where short_url = $1;

-- name: GetUrls :many
select * from urls order by id;

-- name: ListUrls :many
select * from urls order by id limit $1 offset $2;

-- name: DeleteUrlById :exec
delete from urls where id = $1;

-- name: DeleteUrlByShortUrl :exec
delete from urls where short_url = $1;
