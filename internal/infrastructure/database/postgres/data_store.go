package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore struct.
type DataStore struct {
	db initializers.DB
}

// NewDataStore creates and returns a new instance of DataStore.
func NewDataStore(db initializers.DbOps) database.DataStore {
	return &DataStore{db: db}
}

// WithTransaction method.
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	db, ok := d.db.(initializers.DbOps)
	if !ok {
		return errlst.ErrTypeAssertInTransaction
	}

	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return errlst.ErrBeginTransaction
	}

	defer func() {
		if err != nil {
			err = tx.RollBack(ctx)
			if err != nil {
				log.Printf("can not rollback actions in db: %v", err)
			}
		}
	}()

	transactionalDB := &DataStore{db: tx}

	err = transactionFn(transactionalDB)
	if err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return errlst.ErrCommitTx
	}

	return nil
}
