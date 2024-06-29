package sqlite

/*
* SQLite Database Instance
* New, Exec, Query
 */

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func New(databasePath string) error {
	var err error
	db, err = sql.Open("sqlite3", databasePath)
	if err != nil {
		return err
	}
	return nil
}

func Close() error {

	return db.Close()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

func GetDatabaseInstance() *sql.DB {
	return db
}
