-- +goose Up
CREATE TYPE tipe_materi AS ENUM ('kuis', 'materi', 'tugas');

CREATE TABLE pengguna (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(128) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    telepon VARCHAR(32) NOT NULL,
    password VARCHAR(255) NOT NULL,
    dibuat TIMESTAMP DEFAULT NOW()
);

CREATE TABLE kelas (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(64) NOT NULL,
    subjek VARCHAR(64) NOT NULL,
    pengajar INTEGER REFERENCES pengguna (id) NOT NULL,
    kode VARCHAR(6) NOT NULL UNIQUE,
    dibuat TIMESTAMP DEFAULT NOW()
);

CREATE TABLE murid (
    id SERIAL PRIMARY KEY,
    id_pengguna INTEGER REFERENCES pengguna (id) NOT NULL,
    id_kelas INTEGER REFERENCES kelas (id) NOT NULL,
    bergabung TIMESTAMP DEFAULT NOW()
);

CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(128) NOT NULL,
    deskripsi VARCHAR(255) NOT NULL,
    id_kelas INTEGER REFERENCES kelas (id) NOT NULL,
    tipe tipe_materi NOT NULL
);

-- +goose Down
DROP TABLE murid;
DROP TABLE post;
DROP TABLE kelas;
DROP TABLE pengguna;
DROP TYPE tipe_materi;
