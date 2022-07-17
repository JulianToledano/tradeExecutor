package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDb struct {
	db *sql.DB
}

func NewSqlite(dbFile string) (db *SqliteDb, err error) {
	sqliteDb, err := sql.Open("sqlite3", dbFile)
	return &SqliteDb{sqliteDb}, err
}

func (s SqliteDb) EnsureTables() (err error) {
	orders := `
	CREATE TABLE IF NOT EXISTS orders(
	id VARCHAR(64) PRIMARY KEY,
	size REAL NOT NULL,
	price REAL NOT NULL,
	symbol VARCHAR(64) NOT NULL,
	buy INTEGER NOT NULL,
	date DATE NOT NULL
	);`
	_, err = s.db.Exec(orders)
	if err != nil {
		return
	}

	trades := `CREATE TABLE IF NOT EXISTS trades(
	amount REAL NOT NULL,
	price REAL NOT NULL,
	buy INTEGER NOT NULL,
	orderId INT NOT NULL,
    date DATE NOT NULL,
	FOREIGN KEY(orderId) REFERENCES orders(id)
	);`
	_, err = s.db.Exec(trades)
	return
}

func (s SqliteDb) Insert(query string, args ...any) (err error) {
	_, err = s.db.Exec(query, args...)
	return
}

func (s SqliteDb) List(query string, args ...any) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}
