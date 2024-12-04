package database

import (
	"context"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book"
)

type Transaction func(db DataStore) error

// DataStore interface keeps all need methods for repo layer.
type DataStore interface {
	WithTransaction(ctx context.Context, transaction Transaction) error
	AuthorRepo() author.Repository
	BooksRepo() book.Repository
}
