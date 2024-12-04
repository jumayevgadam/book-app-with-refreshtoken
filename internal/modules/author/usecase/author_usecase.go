package usecase

import (
	"context"

	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
)

var (
	_ author.UseCase = (*AuthorUseCase)(nil)
)

// UseCase struct.
type AuthorUseCase struct {
	repo database.DataStore
}

// NewUseCase func.
func NewAuthorUseCase(repo database.DataStore) *AuthorUseCase {
	return &AuthorUseCase{repo: repo}
}

// CreateAuthor Usecase.
func (u *AuthorUseCase) CreateAuthor(ctx context.Context, request authorModel.Request) (int, error) {
	authorID, err := u.repo.AuthorRepo().CreateAuthor(ctx, request.ToPsqlDBStorage())
	if err != nil {
		return -1, errlst.ErrBadRequest
	}

	return authorID, nil
}
