package author

import (
	"context"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/abstract"
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
)

// Repository interface for authors.
type Repository interface {
	CreateAuthor(ctx context.Context, data authorModel.Response) (int, error)
	GetAuthor(ctx context.Context, authorID int) (*authorModel.AuthorData, error)
	CountAuthors(ctx context.Context) (int, error)
	ListAuthors(ctx context.Context, paginationData abstract.PaginationData) ([]*authorModel.AuthorData, error)
}
