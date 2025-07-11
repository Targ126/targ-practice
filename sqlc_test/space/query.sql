-- name: GetById :one
-- -- timeout :5s
select * from space where id=sqlc.arg('id') limit 1;


-- name: GetByName :one
-- -- timeout :5s
select * from space where name=sqlc.arg('name') limit 1;


-- name: UpdateName1 :exec
-- -- timeout :5s
update space set name=sqlc.arg('name'), updated_at=now() where id=sqlc.arg('id');


-- name: UpdateName2 :one
-- -- timeout :5s
update space set name=sqlc.arg('name'), updated_at=now() where id=sqlc.arg('id') returning *;


-- name: GetByIdWithCache :one
-- -- timeout :5s
-- -- cache: 10s
select * from space where id=sqlc.arg('id') limit 1;



-- name: UpdateName3 :one
-- -- timeout :5s
-- -- invalidate: [GetByIdWithCache]
update space set name=sqlc.arg('name'), updated_at=now() where id=sqlc.arg('id') returning *;

