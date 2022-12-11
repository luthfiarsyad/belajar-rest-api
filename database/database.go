package database

import (
	"belajar-rest-api/helper"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/belajar_restful_api")
	helper.PanicIfError(err)
	return db
}
