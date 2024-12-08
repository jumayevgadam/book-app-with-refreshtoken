package author

import (
	"context"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/abstract"
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
)

// UseCase interface for authors.
type UseCase interface {
	CreateAuthor(ctx context.Context, request authorModel.Request) (int, error)
	GetAuthor(ctx context.Context, authorID int) (*authorModel.Author, error)
	ListAuthors(ctx context.Context, pagination abstract.Pagination) (abstract.PaginatedResponse[*authorModel.Author], error)
}
