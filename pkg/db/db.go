package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password123"
	dbName := "todo"
	dbURI := "(127.0.0.1:3306)"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp"+dbURI+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}