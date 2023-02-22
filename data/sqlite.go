package data

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var Sqlite3 *sql.DB

func SqliteMustInit() {
	var err error
	Sqlite3, err = sql.Open("sqlite3", os.Getenv("SQLITE_DBFILE"))
	if err != nil {
		panic("sqlite initiate error")
	}

	Sqlite3.SetMaxOpenConns(1)
}
