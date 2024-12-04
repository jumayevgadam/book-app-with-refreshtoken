package repository

import (
	"context"

	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
)

var (
	_ author.Repository = (*AuthorRepository)(nil)
)

// AuthorRepository struct.
type AuthorRepository struct {
	psqlDB initializers.DB
}

// NewAuthorRepository func.
func NewAuthorRepository(psqlDB initializers.DB) *AuthorRepository {
	return &AuthorRepository{psqlDB: psqlDB}
}

// AuthorRepository repo func.
func (r *AuthorRepository) CreateAuthor(ctx context.Context, authorData authorModel.Response) (int, error) {
	var authorID int

	err := r.psqlDB.QueryRow(
		ctx,
		createAuthorQuery,
		authorData.UserName,
		authorData.Email,
		authorData.PhoneNumber,
		authorData.Password,
		authorData.Bio,
		authorData.Avatar,
	).Scan(&authorID)
	if err != nil {
		return -1, errlst.ErrBadQueryParams
	}

	return authorID, nil
}
