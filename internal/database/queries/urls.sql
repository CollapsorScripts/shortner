-- name: CreateUrl :one
insert into urls (original_url, short_url) values ($1, $2) returning *;

-- name: GetOriginalUrlByShortUrl :one
select original_url from urls where short_url = $1;

-- name: GetOriginalUrlById :one
select original_url from urls where id = $1;
