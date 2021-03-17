package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "username:password@/dbname")
	if err != nil {
		panic(err.Error())
	}

	return db
}