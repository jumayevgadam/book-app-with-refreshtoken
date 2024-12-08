package initializers

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
)

var _ DB = (*Database)(nil)

// Querier interface for using pgxscany.
type Querier interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

// DB interface handles needed DB method
type DB interface {
	Querier
	Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error
}

// DbOps interface with transaction.
type DbOps interface {
	DB
	Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error)
	Close() error
}

// Get method implements DB interface.
func (d *Database) Get(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db, dest, query, args...)
}

// Select method implements DB interface.
func (d *Database) Select(ctx context.Context, db Querier, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db, dest, query, args...)
}

// QueryRow method implements DB interface.
func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Db.QueryRow(ctx, query, args...)
}

// Query method implements DB interface.
func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return d.Db.Query(ctx, query, args...)
}

// Exec method implements DB interface.
func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return d.Db.Exec(ctx, query, args...)
}

// Begin starts a new transaction.
func (d *Database) Begin(ctx context.Context, txOpts pgx.TxOptions) (TxOps, error) {
	if d == nil {
		return nil, errlist.ErrBeginTransaction
	}

	c, err := d.Db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("acquire connection: %w", errlist.ErrBeginTransaction)
	}

	tx, err := d.Db.BeginTx(ctx, txOpts)
	if err != nil {
		c.Release()
	}

	return &Transaction{Tx: tx}, nil
}

// Close closes the database connection pool.
func (d *Database) Close() error {
	d.Db.Close()
	return nil
}
