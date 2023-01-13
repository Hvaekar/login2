package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB
var ErrorDB error

// Открытие БД
func OpenDB(dsn string) () {
	DB, ErrorDB = sql.Open("mysql", dsn)
	if ErrorDB != nil {
		log.Fatal(ErrorDB)
	}
	if ErrorDB = DB.Ping(); ErrorDB != nil {
		log.Fatal(ErrorDB)
	}
}

func CloseDB(DB *sql.DB) {
	DB.Close()
}
