package tradeExecutor

import "database/sql"

type DataBase interface {
	EnsureTables() error
	Insert(query string, args ...any) error
	List(query string, args ...any) (*sql.Rows, error)
}
