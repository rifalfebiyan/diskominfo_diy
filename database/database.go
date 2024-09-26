package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	dsn := "root@tcp(localhost:3306)/buku_tamu"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil { // Perbaikan 'nill' menjadi 'nil'
		panic(err)
	}
	return db
}
