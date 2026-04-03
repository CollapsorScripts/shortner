-- name: CreateStatistics :one
insert into statistics (url_id) values ($1) returning *;

-- name: UpdateLastAccessedById :one
update statistics set last_accessed = now() where id = $1 returning *;

-- name: UpdateLastAccessedByUrlId :one
update statistics set last_accessed = now() where url_id = $1 returning *;

-- name: IncrementClicksCountByUrlId :one
update statistics set access_count = access_count + 1 where url_id = $1 returning *;

-- name: IncrementClicksCountById :one
update statistics set access_count = access_count + 1 where id = $1 returning *;

-- name: GetStatistics :many
select * from statistics order by id;

-- name: GetStatisticById :one
select * from statistics where id = $1 order by id;

-- name: GetStatisticsByUrlId :one
select * from statistics where url_id = $1 order by id;

-- name: ListStatistics :many
select * from statistics order by id limit $1 offset $2;
