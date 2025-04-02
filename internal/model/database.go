package model

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DB interface {
	Exec(query string, args ...any) (Result, error)
}
