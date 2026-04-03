-- name: CreateUserAgent :one
insert into user_agents (fingerprint_id, agent) values ($1, $2) returning *;

-- name: GetUserAgents :many
select * from user_agents order by id;

-- name: GetUserAgentById :one
select * from user_agents where id = $1;

-- name: GetUserAgentsByFingerprintId :many
select * from user_agents where fingerprint_id = $1;

-- name: UpdateUserAgentLastAccessedById :one
update user_agents set last_accessed = now() where id = $1 returning *;
