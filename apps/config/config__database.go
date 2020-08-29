package config

import (
	"ZebraX/apps/handle"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var databases *sql.DB

//Getdatabase is init database connection
func Getdatabase() *sql.DB {
	if databases == nil {
		db, err := sql.Open(Getenv("Database"), Getenv("Connection"))
		err = db.Ping()
		handle.ErrorHandle(err)
		databases = db
	}
	return databases
}
