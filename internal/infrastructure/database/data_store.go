package database

import "context"

type Transaction func(db DataStore) error

// DataStore interface keeps all need methods for repo layer.
type DataStore interface {
	WithTransaction(ctx context.Context, transaction Transaction) error
}
