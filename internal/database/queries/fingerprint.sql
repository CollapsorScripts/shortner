-- name: CreateFingerPrint :one
insert into fingerprint (statistics_id, ip) values ($1, $2) returning *;

-- name: GetFingerPrints :many
select * from fingerprint order by id;

-- name: ListFingerPrint :many
select * from fingerprint order by id limit $1 offset $2;

-- name: ListFingerPrintByStatisticsId :many
select * from fingerprint where statistics_id = $1 order by id limit $2 offset $3;

-- name: GetFingerPrintByIp :one
select * from fingerprint where ip = $1 order by id;

-- name: GetFullFingerPrintById :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprint fp join user_agents ua on fp.id = ua.fingerprint_id where fp.id = $1;

-- name: GetFullFingerPrintByStatisticsId :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprint fp join user_agents ua on fp.id = ua.fingerprint_id where fp.statistics_id = $1;

-- name: GetFullFingerPrintByIp :one
select
    sqlc.embed(fp),
    sqlc.embed(ua)
from fingerprint fp join user_agents ua on fp.id = ua.fingerprint_id where fp.ip = $1;
