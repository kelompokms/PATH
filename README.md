# PATH

Personal Adaptive Teaching Hub (PATH)

## Setup

### Kebutuhan Menjalankan

- Docker

### Cara Menjalankan

```bash
docker compose up -d --build
docker exec -it path-api-1 sh
goose run ./internal/db/migrations postgres $DATABASE_URL up
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
