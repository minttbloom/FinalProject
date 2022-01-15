package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open("mysql", "user:password@/db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
