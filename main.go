package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	"log"
)

var conn *sql.DB

const serverAddress = "0.0.0.0:8080"

func main() {
	cfg := mysql.Config{
		User:                 ("root"),
		Passwd:               ("secret"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "simplebank",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	// Get a database handle.
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewSever(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
