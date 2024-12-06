package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {

	dns := "root:@tcp(127.0.0.1:3306)/karyawan"
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db

}
