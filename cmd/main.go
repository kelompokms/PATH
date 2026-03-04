package main

import (
	"database/sql"
	"log"
	"path_project/internal/handler"
	"path_project/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB, err := sql.Open("mysql", "rin:blank@/path?parseTime=true")
	if err != nil {
		log.Fatalln("Failed to open database:", err)
	}

	queries := db.New(DB)

	server := handler.NewServer(queries)

	server.Start(":3000")
}
