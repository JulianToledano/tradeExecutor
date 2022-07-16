package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: renaming db.db :/
type SqliteDb struct {
	db *sql.DB
}

func NewSqlite(dbFile string) (db *SqliteDb, err error) {
	sqliteDb, err := sql.Open("sqlite3", dbFile)
	return &SqliteDb{sqliteDb}, err
}

func (db SqliteDb) EnsureTables() (err error) {
	create := `
	CREATE TABLE IF NOT EXISTS orders(
	size REAL NOT NULL,
	price REAL NOT NULL,
	symbol VARCHAR(64) NOT NULL,
	buy INTEGER NOT NULL,
	date DATE NOT NULL
	);`
	_, err = db.db.Exec(create)
	return
}

func (db SqliteDb) Insert(query string, args ...any) (err error) {
	_, err = db.db.Exec(query, args...)
	return
}
