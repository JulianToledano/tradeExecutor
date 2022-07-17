package db

import "database/sql"

type MockDb struct{}

func NewMockDb() *MockDb                                           { return &MockDb{} }
func (m MockDb) EnsureTables() error                               { return nil }
func (m MockDb) Insert(query string, args ...any) error            { return nil }
func (m MockDb) List(query string, args ...any) (*sql.Rows, error) { return nil, nil }
