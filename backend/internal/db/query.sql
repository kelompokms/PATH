-- name: ListMurid :many
select * from murid where kode_kelas = $1;

-- name: ListPost :many
select * from post where kode_kelas = $1;

-- name: ListKelas :many
SELECT kelas.id, kelas.nama as nama_kelas, subjek, pengguna.nama as nama_pengguna, kode, kelas.dibuat FROM kelas join pengguna on kelas.pengajar = pengguna.id WHERE kelas.pengajar = $1
UNION
SELECT kelas.id, kelas.nama as nama_kelas, subjek, pengguna.nama as nama_pengguna, kode, kelas.dibuat FROM kelas JOIN murid ON murid.kode_kelas = kelas.kode join pengguna on kelas.pengajar = pengguna.id
WHERE murid.id_pengguna = $1;

-- name: GetKelas :one
select k.id, k.nama as nama_kelas, subjek, p.nama as nama_pengajar, kode, k.dibuat as dibuat from kelas k
join pengguna p on p.id = k.pengajar where kode = $1;

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
