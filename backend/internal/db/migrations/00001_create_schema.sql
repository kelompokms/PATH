-- +goose Up
CREATE TYPE tipe_materi AS ENUM ('kuis', 'materi', 'tugas');
CREATE TYPE tipe_kelamin AS ENUM ('l', 'p');

CREATE TABLE pengguna (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(128) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    jenis_kelamin tipe_kelamin NOT NULL,
    telepon VARCHAR(32) NOT NULL,
    password VARCHAR(255) NOT NULL,
    dibuat TIMESTAMP DEFAULT NOW()
);

CREATE TABLE kelas (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(64) NOT NULL,
    bagian VARCHAR(64),
    deskripsi VARCHAR(320),
    pengajar INTEGER REFERENCES pengguna (id) NOT NULL,
    kode VARCHAR(6) NOT NULL UNIQUE,
    dibuat TIMESTAMP DEFAULT NOW()
);

CREATE TABLE murid (
    id SERIAL PRIMARY KEY,
    id_pengguna INTEGER REFERENCES pengguna (id) NOT NULL,
    kode_kelas VARCHAR(6) REFERENCES kelas (kode) NOT NULL,
    bergabung TIMESTAMP DEFAULT NOW()
);

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(128) NOT NULL,
    deskripsi VARCHAR(255) NOT NULL,
    kode_kelas VARCHAR(6) REFERENCES kelas (kode) NOT NULL,
    tipe tipe_materi NOT NULL
);

-- +goose Down
DROP TABLE murid;
DROP TABLE post;
DROP TABLE kelas;
DROP TABLE pengguna;
DROP TYPE tipe_materi;
DROP TYPE tipe_kelamin;
