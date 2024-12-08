package repository

import (
	"context"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/abstract"
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
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
		authorData.Password,
		authorData.Bio,
		authorData.Avatar,
	).Scan(&authorID)
	if err != nil {
		return -1, errlist.ParseSQLErrors(err)
	}

	return authorID, nil
}

// GetAuthor repo.
func (r *AuthorRepository) GetAuthor(ctx context.Context, authorID int) (*authorModel.AuthorData, error) {
	var author authorModel.AuthorData

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&author,
		getAuthorQuery,
		authorID,
	)
	if err != nil {
		return nil, errlist.ParseSQLErrors(err)
	}

	return &author, nil
}

// CountAuthors repo.
func (r *AuthorRepository) CountAuthors(ctx context.Context) (int, error) {
	var totalCount int

	err := r.psqlDB.Get(
		ctx,
		r.psqlDB,
		&totalCount,
		countAuthorsQuery,
	)
	if err != nil {
		return -1, errlist.ParseSQLErrors(err)
	}

	return totalCount, nil
}

// ListAuthors repo.
func (r *AuthorRepository) ListAuthors(ctx context.Context, pgData abstract.PaginationData) ([]*authorModel.AuthorData, error) {
	var authorList []*authorModel.AuthorData
	offset := (pgData.CurrentPage - 1) * pgData.Limit

	err := r.psqlDB.Select(
		ctx,
		r.psqlDB,
		&authorList,
		listAuthorQuery,
		offset,
		pgData.Limit,
	)
	if err != nil {
		return nil, errlist.ParseSQLErrors(err)
	}

	return authorList, nil
}
