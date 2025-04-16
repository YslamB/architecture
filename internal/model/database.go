package model

import (
	"context"
	"database/sql"
)

type CommandTag interface {
	RowsAffected() int64
}

type DB interface {
	Exec(ctx context.Context, query string, args ...any) (CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...any) *sql.Row
}
