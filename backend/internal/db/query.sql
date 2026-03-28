-- name: ListMurid :many
select * from murid where kode_kelas = $1;

-- name: ListPost :many
select * from post where kode_kelas = $1;

-- name: ListKelas :many
SELECT kelas.id, nama, subjek, pengajar, kode, dibuat FROM kelas WHERE kelas.pengajar = $1
UNION
SELECT kelas.id, nama, subjek, pengajar, kode, dibuat FROM kelas JOIN murid ON murid.kode_kelas = kelas.kode
WHERE murid.id_pengguna = $1;

-- name: GetKelas :one
select * from kelas where kode = $1;

-- name: CreateKelas :exec
insert into kelas (nama, subjek, pengajar, kode) values ($1, $2, $3, $4);

-- name: CreatePost :exec
insert into post (nama, deskripsi, kode_kelas, tipe) values ($1, $2, $3, $4);

-- name: ListPengguna :many
select * from pengguna;

-- name: GetPengguna :one
select * from pengguna where id = $1;

-- name: ValidatePenggunaID :one
select id from pengguna where id = $1;

-- name: CheckPengguna :one
select id, password from pengguna where email = $1;

-- name: CreatePengguna :one
insert into pengguna(nama, email, telepon, password) values ($1, $2, $3, $4) returning id;
