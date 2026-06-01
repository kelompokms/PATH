-- name: ListMurid :many
select id_pengguna, nama, email from murid join pengguna on id_pengguna = pengguna.id where kode_kelas = $1;

-- name: ListPost :many
select id, nama, deskripsi, tipe, tenggat, dibuat, file from post where kode_kelas = $1 order by dibuat desc;

-- name: ListPostTugas :many
select id, nama, deskripsi, tenggat, dibuat, file from post where kode_kelas = $1 and tipe = 'tugas' order by dibuat desc;

-- name: GetPost :one
select id, nama, deskripsi, tipe, tenggat, dibuat, file from post where kode_kelas = $1 and id = $2;

-- name: CreatePost :exec
insert into post (nama, deskripsi, kode_kelas, tenggat, tipe, file) values ($1, $2, $3, $4, $5, $6);


-- name: ListKelas :many
select kelas.id, kelas.nama as nama_kelas, bagian, pengguna.nama as nama_pengguna, kode, kelas.dibuat from kelas join pengguna on kelas.pengajar = pengguna.id WHERE kelas.pengajar = $1
union
select kelas.id, kelas.nama as nama_kelas, bagian, pengguna.nama as nama_pengguna, kode, kelas.dibuat from kelas join murid on murid.kode_kelas = kelas.kode join pengguna on kelas.pengajar = pengguna.id
where murid.id_pengguna = $1;

-- name: GetKelas :one
select k.nama as nama_kelas, k.kode, bagian, p.nama as nama_pengajar, $2 = p.id as is_pengajar from kelas k
join pengguna p on p.id = k.pengajar where kode = $1;

-- name: CreateKelas :exec
insert into kelas (nama, bagian, pengajar, kode) values ($1, $2, $3, $4);

-- name: GetKelasCode :many
select * from kelas where kode = $1;

-- name: CheckPenggunaInKelas :one
select murid.id from murid where id_pengguna = $1 and kode_kelas = $2;

-- name: JoinKelas :exec
insert into murid (id_pengguna, kode_kelas) values ($1, $2);


-- name: ListPengguna :many
select * from pengguna;

-- name: GetPengguna :one
select id, nama, email, jenis_kelamin, telepon from pengguna where id = $1;

-- name: ValidatePenggunaID :one
select id from pengguna where id = $1;

-- name: CheckPengguna :one
select id, password from pengguna where email = $1;

-- name: CreatePengguna :one
insert into pengguna(nama, email, telepon, password, jenis_kelamin) values ($1, $2, $3, $4, $5) returning id;
