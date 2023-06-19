package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

var db *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	cfg := mysql.Config{
		User:                 ("root"),
		Passwd:               ("secret"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "simplebank",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(db)
	os.Exit(m.Run())
}
