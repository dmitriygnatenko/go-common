package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TxDB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
}

type TxManager struct {
	db TxDB
}

func NewTransactionManager(db TxDB) *TxManager {
	return &TxManager{
		db: db,
	}
}

func (s *TxManager) transaction(ctx context.Context, opts sql.TxOptions, fn func(ctx context.Context) error) (err error) {
	tx, ok := ctx.Value(TxKey{}).(*sqlx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = s.db.BeginTx(ctx, &opts)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	ctx = context.WithValue(ctx, TxKey{}, tx)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}

		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				err = fmt.Errorf("transaction rollback: %w", errRollback)
			}

			return
		}

		if nil == err {
			err = tx.Commit()
			if err != nil {
				err = fmt.Errorf("transaction commit: %w", err)
			}
		}
	}()

	if err = fn(ctx); err != nil {
		err = fmt.Errorf("failed executing code inside transaction: %w", err)
	}

	return err
}

func (s *TxManager) ReadCommitted(ctx context.Context, f func(ctx context.Context) error) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelReadCommitted}
	return s.transaction(ctx, txOpts, f)
}

func (s *TxManager) RepeatableRead(ctx context.Context, f func(ctx context.Context) error) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelRepeatableRead}
	return s.transaction(ctx, txOpts, f)
}

func (s *TxManager) Serializable(ctx context.Context, numAttempts int, f func(ctx context.Context) error) error {
	txOpts := sql.TxOptions{Isolation: sql.LevelSerializable}

	for i := 0; i < numAttempts; i++ {
		err := s.transaction(ctx, txOpts, f)
		if err != nil {
			continue
		}

		return nil
	}

	return fmt.Errorf("serialization error after %d attempts", numAttempts)
}
