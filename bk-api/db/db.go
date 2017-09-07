package db

import (
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"bk-api/utils"
)

var readOnlyDB *sql.DB
var db *sql.DB

func init() {
	var err error
	//db, err = sql.Open("sqlite3", "/var/lib/mageddo/bookmarks-node/db/bookmarks.db?mode=ro")

	readOnlyDB, err = sql.Open("sqlite3", utils.GetConfig().DatabaseURL)
	readOnlyDB.SetMaxOpenConns(5)
	if (err != nil) {
		panic(err)
	}

	db, err = sql.Open("sqlite3", utils.GetConfig().DatabaseURL)
	db.SetMaxOpenConns(1)
	if (err != nil) {
		panic(err)
	}
}

// Read only connection
func GetROConn() *sql.DB {
	return readOnlyDB
}

func GetConn() *sql.DB {
	if utils.IsTestProfile() {
		return readOnlyDB
	}
	return db
}
