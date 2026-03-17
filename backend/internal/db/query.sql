-- name: ListMurid :many
select * from murid where id_kelas = $1;

-- name: ListPost :many
select * from post where id_kelas = $1;

-- name: ListKelas :many
select * from kelas;

-- name: ListPengguna :many
select * from pengguna;
