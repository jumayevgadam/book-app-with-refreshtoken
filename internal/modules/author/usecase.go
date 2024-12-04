package author

import (
	"context"

	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
)

// UseCase interface for authors.
type UseCase interface {
	CreateAuthor(ctx context.Context, request authorModel.Request) (int, error)
}
