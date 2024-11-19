package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TxKey struct{}

type DB struct {
	db *sqlx.DB
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

	var source string

	switch c.driver {
	case "mysql":
		source = c.username +
			":" + c.password +
			"@tcp(" + c.host + ":" + strconv.Itoa(int(c.port)) + ")/" +
			c.dbname +
			"?parseTime=true"
	case "postgres":
		source = "user=" + c.username +
			" password=" + c.password +
			" dbname=/" + c.dbname +
			" host=" + c.host +
			" port=" + strconv.Itoa(int(c.port))
	}

	sqlConn, err := sql.Open(c.driver, source)
	if err != nil {
		return nil, fmt.Errorf("open DB connection error: %w", err)
	}

	db := sqlx.NewDb(sqlConn, c.driver)

	if c.maxOpenConns > 0 {
		db.SetMaxOpenConns(int(c.maxOpenConns))
	}

	if c.maxOpenConns > 0 {
		db.SetMaxIdleConns(int(c.maxIdleConns))
	}

	if c.maxOpenConnLifetime != nil {
		db.SetConnMaxLifetime(*c.maxOpenConnLifetime)
	}

	if c.maxIdleConnLifetime != nil {
		db.SetConnMaxIdleTime(*c.maxIdleConnLifetime)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB ping error: %w", err)
	}

	return &DB{db: db}, nil
}

func (s *DB) Close() error {
	return s.db.Close()
}

func (s *DB) Ping() error {
	return s.db.Ping()
}

func (s *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return s.db.BeginTxx(ctx, opts)
}

func (s *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx, ok := ctx.Value(TxKey{}).(*sqlx.Tx)
	if ok {
		return tx.SelectContext(ctx, dest, query, args...)
	}

	return s.db.SelectContext(ctx, dest, query, args...)
}

func (s *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx, ok := ctx.Value(TxKey{}).(*sqlx.Tx)
	if ok {
		return tx.GetContext(ctx, dest, query, args...)
	}

	return s.db.GetContext(ctx, dest, query, args...)
}

func (s *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	tx, ok := ctx.Value(TxKey{}).(*sqlx.Tx)
	if ok {
		return tx.ExecContext(ctx, query, args...)
	}

	return s.db.ExecContext(ctx, query, args...)
}

func (s *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx, ok := ctx.Value(TxKey{}).(*sqlx.Tx)
	if ok {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return s.db.QueryRowContext(ctx, query, args...)
}
