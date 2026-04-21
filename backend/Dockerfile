FROM golang:1.26-alpine

WORKDIR /usr/local/bin/
RUN wget https://github.com/pressly/goose/releases/download/v3.27.0/goose_linux_x86_64 -o goose
RUN chmod +x /usr/local/bin/goose

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN goose -dir ./internal/db/migrations postgres "postgresql://wave:123123@path-db-1:5433/path?sslmode=disable" status
RUN goose --version
RUN go build -o ./path_api
