package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseType = "mysql"
var DatabaseUrl = "root:@tcp(127.0.0.1:3306)/chulgeun_gil_planner"
var Db *sql.DB

func ConnectDb() {
	db, err := sql.Open(DatabaseType, DatabaseUrl)
	if err != nil {
		panic(err)
	}
	Db = db
}
