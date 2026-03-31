# PATH Backend

Personal Adaptive Teaching Hub (PATH)

## Prerequisites

- Go 1.26
- [Goose](https://pressly.github.io/goose/installation/) (migration tool)
- Postgresql 18

(opsional)

- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html) (untuk compile sql ke code)

## Cara menjalankan

1. clone repositori ini

```bash
git clone https://github.com/kelompokms/path.git
```

2. Download dependencies

```bash
cd path
go mod download
```

3. Jalankan migration

```bash
goose -dir ./internal/db/migrations postgres "user=USER password=PASSWORD dbname=DB sslmode=disable" up
```

4. Jalankan

```bash
DATABASE_URL="<ISIKAN URL DATABASE AKTIF>" go run cmd/main.go
```

## Cara development

1. Jika mengubah `internal/db/query.sql` atau `internal/db/migrations/`, jalankan `sqlc generate`
