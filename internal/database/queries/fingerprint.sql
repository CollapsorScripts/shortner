-- name: CreateFingerPrint :one
insert into fingerprints (statistics_id, ip) values ($1, $2) returning *;

-- name: GetFingerPrints :many
select * from fingerprints order by id;

-- name: ListFingerPrint :many
select * from fingerprints limit $1 offset $2 order by id;

-- name: ListFingerPrintByStatisticsId :many
select * from fingerprints where statistics_id = $1 limit $2 offset $3 order by id;

-- name: GetFingerPrintByIp :one
select * from fingerprints where ip = $1 order by id;

-- name: GetFullFingerPrintById :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprints fp join user_agents ua on fp.id = ua.fingerprint_id where fp.id = $1;

-- name: GetFullFingerPrintByStatisticsId :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprints fp join user_agents ua on fp.id = ua.fingerprint_id where fp.statistics_id = $1;

-- name: GetFullFingerPrintByIp :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprints fp join user_agents ua on fp.id = ua.fingerprint_id where fp.ip = $1;
