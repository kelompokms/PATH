-- name: ListMurid :many
select * from murid where id_kelas = $1;

-- name: ListPost :many
select * from post where id_kelas = $1;

-- name: ListKelas :many
select * from kelas;

-- name: ListPengguna :many
select * from pengguna;

-- name: GetPengguna :one
select * from pengguna where id = $1;

-- name: CheckPengguna :one
select id, password from pengguna where email = $1;

-- name: CreatePengguna :one
insert into pengguna(nama, email, telepon, password) values ($1, $2, $3, $4) returning id;
