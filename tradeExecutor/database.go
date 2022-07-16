package tradeExecutor

type DataBase interface {
	EnsureTables() error
	Insert(query string, args ...any) error
}
