package dbops

import (
	"database/sql"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(47.93.48.102:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
