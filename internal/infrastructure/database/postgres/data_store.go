package postgres

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	authorRepository "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author/repository"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book"
	bookRepository "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book/repository"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
)

var _ database.DataStore = (*DataStore)(nil)

// DataStore struct.
type DataStore struct {
	db         initializers.DB
	authorInit sync.Once
	author     author.Repository
	bookInit   sync.Once
	book       book.Repository
}

// NewDataStore creates and returns a new instance of DataStore.
func NewDataStore(db initializers.DbOps) database.DataStore {
	return &DataStore{db: db}
}

func (d *DataStore) AuthorRepo() author.Repository {
	d.authorInit.Do(func() {
		d.author = authorRepository.NewAuthorRepository(d.db)
	})

	return d.author
}

func (d *DataStore) BooksRepo() book.Repository {
	d.bookInit.Do(func() {
		d.book = bookRepository.NewBookRepository(d.db)
	})

	return d.book
}

// WithTransaction method.
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	db, ok := d.db.(initializers.DbOps)
	if !ok {
		return errlist.ErrTypeAssertInTransaction
	}

	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return errlist.ErrBeginTransaction
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
		return errlist.ParseErrors(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return errlist.ErrCommitTx
	}

	return nil
}
