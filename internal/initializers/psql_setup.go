package initializers

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
)

// Database struct keeps pgxpool.
type Database struct {
	Db *pgxpool.Pool
}

// GetDBConnection with needed params.
func GetDBConnection(ctx context.Context, cfgs config.Config) (*Database, error) {
	db, err := pgxpool.New(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfgs.Postgres.User,
		cfgs.Postgres.Password,
		cfgs.Postgres.Host,
		cfgs.Postgres.Port,
		cfgs.Postgres.Name,
		cfgs.Postgres.SslMode,
	))
	if err != nil {
		return nil, errlist.ErrDBConnection
	}

	// give ping, if error occured return error.
	err = db.Ping(ctx)
	if err != nil {
		return nil, errlist.ErrDBPing
	}

	return &Database{Db: db}, nil
}
