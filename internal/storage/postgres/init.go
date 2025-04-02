package postgres

import (
	"architecture/internal/model"
	"database/sql"
)

type PsqlDB struct {
	db *sql.DB
}

type Database struct {
	PsqlDB
}

func (db *PsqlDB) Exec(query string, args ...interface{}) (model.Result, error) {
	return db.db.Exec(query, args...)
}

// func (db *PsqlDB) Query(query string, args ...any) (*model.Rows, error) {
// 	return db.db.Query(query, args...)
// }

// func (db *PsqlDB) QueryRow(query string, args ...any) *model.Row {
// 	return db.db.QueryRow(query, args...)
// }

func Init(connString string) (model.DB, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Database{
		PsqlDB: PsqlDB{db: db},
	}, nil
}
