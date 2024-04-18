package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type TxKey struct{}

type DB struct {
	db *sql.DB
}

func NewDB(c Config) (*DB, error) {
	if len(c.username) == 0 {
		return nil, errors.New("empty username")
	}

	if len(c.password) == 0 {
		return nil, errors.New("empty password")
	}

	if len(c.dbname) == 0 {
		return nil, errors.New("empty database name")
	}

	if len(c.driver) == 0 {
		c.driver = defaultDriver
	}

	if len(c.host) == 0 {
		c.host = defaultHost
	}

	if c.port == 0 {
		c.port = defaultPort
	}

	if len(c.sslMode) == 0 {
		c.sslMode = defaultSslMode
	}

	source := "user=" + c.username +
		" password=" + c.password +
		" dbname=" + c.dbname +
		" host=" + c.host +
		" port=" + strconv.Itoa(int(c.port)) +
		" sslmode=" + c.sslMode

	db, err := sql.Open(c.driver, source)
	if err != nil {
		return nil, fmt.Errorf("open DB connection error: %w", err)
	}

	if c.maxOpenConns > 0 {
		db.SetMaxOpenConns(int(c.maxOpenConns))
	}

	if c.maxOpenConns > 0 {
		db.SetMaxIdleConns(int(c.maxIdleConns))
	}

	if c.maxConnLifetime != nil {
		db.SetConnMaxLifetime(*c.maxConnLifetime)
	}

	if c.maxIdleConnLifetime != nil {
		db.SetConnMaxIdleTime(*c.maxIdleConnLifetime)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB ping error: %w", err)
	}

	return &DB{db: db}, nil
}

func (s *DB) Ping() error {
	return s.db.Ping()
}

func (s *DB) Close() error {
	return s.db.Close()
}

func (s *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return s.db.BeginTx(ctx, opts)
}

func (s *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.QueryContext(ctx, query, args...)
	}

	return s.db.QueryContext(ctx, query, args...)
}

func (s *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return s.db.QueryRowContext(ctx, query, args...)
}

func (s *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return tx.ExecContext(ctx, query, args...)
	}

	return s.db.ExecContext(ctx, query, args...)
}
