package mockdb

import (
	"architecture/internal/model"
)

type PsqlDB struct {
}

type Database struct {
	PsqlDB
}
type MockResult struct {
	LastInsertIDValue int64
	RowsAffectedValue int64
}

func (r *MockResult) LastInsertId() (int64, error) {
	return r.LastInsertIDValue, nil
}

func (r *MockResult) RowsAffected() (int64, error) {
	return r.RowsAffectedValue, nil
}

func (db *PsqlDB) Exec(query string, args ...interface{}) (model.Result, error) {
	return &MockResult{}, nil
}

func Init(connString string) (model.DB, error) {

	return &Database{
		PsqlDB: PsqlDB{},
	}, nil
}
