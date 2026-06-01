# PATH

Personal Adaptive Teaching Hub (PATH)

## Setup

### Kebutuhan Menjalankan

- Docker

### Cara Menjalankan

1. Clone repository
```bash
git clone https://github.com/kelompokms/PATH.git
cd PATH
```
2. Rename `.env.example` menjadi `.env`. Lalu isikan (deskripsi di dalam)

3. Pada direktori frontend, rename `.env.example` menjadi `.env`. Lalu isikan (deskripsi di dalam)

4. Jalankan compose
```bash
docker compose up -d --build
```

5. Jalankan migrations
```bash
docker compose up -d migrations
docker exec -it path-migrations-1 sh

# Setelah masuk ke container
goose down
goose up
```

### Kebutuhan Development

- Postgresql
- NodeJS >= 20
- Go >= 1.26
  - [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
  - [goose](https://pressly.github.io/goose/installation/)

### Cara Development

#### Frontend

```bash
cd frontend
pnpm install
pnpm dev
```

#### Backend

```bash
cd backend
go mod download
goose run ./internal/db/migrations postgres $DATABASE_URL up
```
