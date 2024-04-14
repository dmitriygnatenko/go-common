package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type TxKey struct{}

type Config interface {
	Driver() string

	User() string
	Password() string
	Name() string
	Host() string
	Port() string
	SSLMode() string

	MaxOpenConns() int
	MaxIdleConns() int
	MaxConnLifetime() int     // in seconds
	MaxIdleConnLifetime() int // in seconds
}

type DB struct {
	db *sql.DB
}

func NewDB(c Config) (*DB, error) {
	source := "user=" + c.User() +
		" password=" + c.Password() +
		" dbname=" + c.Name() +
		" host=" + c.Host() +
		" port=" + c.Port() +
		" sslmode=" + c.SSLMode()

	db, err := sql.Open(c.Driver(), source)
	if err != nil {
		return nil, fmt.Errorf("open DB connection error: %w", err)
	}

	if c.MaxOpenConns() > 0 {
		db.SetMaxOpenConns(c.MaxOpenConns())
	}

	if c.MaxIdleConns() > 0 {
		db.SetMaxIdleConns(c.MaxIdleConns())
	}

	if c.MaxConnLifetime() > 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(c.MaxConnLifetime()))
	}

	if c.MaxIdleConnLifetime() > 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(c.MaxIdleConnLifetime()))
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB ping error: %w", err)
	}

	return &DB{db: db}, nil
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.db.BeginTx(ctx, opts)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.QueryContext(ctx, query, args...)
	}

	return db.db.QueryContext(ctx, query, args...)
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return db.db.QueryRowContext(ctx, query, args...)
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.ExecContext(ctx, query, args...)
	}

	return db.db.ExecContext(ctx, query, args...)
}
