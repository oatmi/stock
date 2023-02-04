package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Sqlite3 *sql.DB

func SqliteMustInit() {
	var err error
	Sqlite3, err = sql.Open("sqlite3", "/Users/yangtao/Documents/stock_db/stock.sqlite") // TODO read from ENV
	if err != nil {
		panic("sqlite initiate error")
	}

	Sqlite3.SetMaxOpenConns(1)
}
