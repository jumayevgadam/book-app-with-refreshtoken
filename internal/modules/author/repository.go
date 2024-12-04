package author

import (
	"context"

	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
)

// Repository interface for authors.
type Repository interface {
	CreateAuthor(ctx context.Context, data authorModel.Response) (int, error)
}
