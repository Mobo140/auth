package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client is a interface for a work with database
type Client interface {
	DB() DB
	Close() error
}

// DB is a interface for a work with database
type DB interface {
	SQLExecer
	Pinger
	Close()
}

// Query is a wrapper over a query that stores the name of the query and the query itself
type Query struct {
	Name     string
	QueryRow string
}

// /SQLExecer is a combine os QueryExecer and NamedExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer is a interface for a work with named queries using tags in structures
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer is a interface for a work  with common queries
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger is a interface for a checking a database connection
type Pinger interface {
	Ping(ctx context.Context) error
}