package sqlite

/*
* SQLite Database Instance
* New, Exec, Query
 */

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDBInstance struct {
	db *sql.DB
}

func New(databasePath string) (*SQLiteDBInstance, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	return &SQLiteDBInstance{db: db}, nil
}

func (app *SQLiteDBInstance) Close() error {
	return app.db.Close()
}

func (app *SQLiteDBInstance) Exec(query string, args ...interface{}) (sql.Result, error) {
	return app.db.Exec(query, args...)
}

func (app *SQLiteDBInstance) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return app.db.Query(query, args...)
}
