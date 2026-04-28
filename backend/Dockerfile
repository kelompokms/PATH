FROM golang:1.26-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o ./path_api ./cmd/main.go
