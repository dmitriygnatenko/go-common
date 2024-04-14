package db

import (
	"context"
	"database/sql"
	"fmt"
)

type TxDB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type Logger interface {
	ErrorContext(ctx context.Context, msg string, args ...any)
}

type Handler func(ctx context.Context) error

type TxManager struct {
	db TxDB
}

func NewTransactionManager(db TxDB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (tm *TxManager) transaction(ctx context.Context, opts sql.TxOptions, fn Handler) error {
	tx, ok := ctx.Value(TxKey{}).(*sql.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err := tm.db.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("begin transaction  error: %w", err)
	}

	ctx = context.WithValue(ctx, TxKey{}, tx)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = fmt.Errorf("transaction rollback error: %w", errRollback)
			}

			return
		}

		if err == nil {
			err = tx.Commit()
			if err != nil {
				err = fmt.Errorf("transaction commit error: %w", err)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = fmt.Errorf("failed executing code inside transaction: %w", err)
	}

	return err
}

func (tm *TxManager) ReadCommitted(ctx context.Context, f Handler) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelReadCommitted}
	return tm.transaction(ctx, txOpts, f)
}

func (tm *TxManager) RepeatableRead(ctx context.Context, f Handler) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelRepeatableRead}
	return tm.transaction(ctx, txOpts, f)
}

func (tm *TxManager) Serializable(ctx context.Context, numAttempts int, f Handler) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelSerializable}

	for i := 0; i < numAttempts; i++ {
		err := tm.transaction(ctx, txOpts, f)
		if err != nil {
			continue
		}

		return nil
	}

	return fmt.Errorf("serialization error after %d attempts", numAttempts)
}
