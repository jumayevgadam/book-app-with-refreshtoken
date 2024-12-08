package usecase

import (
	"context"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/abstract"
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
	"github.com/samber/lo"
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
		return -1, errlist.ParseErrors(err)
	}

	return authorID, nil
}

// GetAuthor usecase.
// Add redis this function.
func (u *AuthorUseCase) GetAuthor(ctx context.Context, authorID int) (*authorModel.Author, error) {
	author, err := u.repo.AuthorRepo().GetAuthor(ctx, authorID)
	if err != nil {
		return nil, errlist.ParseErrors(err)
	}

	return author.ToServer(), nil
}

// ListAuthors usecase.
func (u *AuthorUseCase) ListAuthors(ctx context.Context, pgReq abstract.Pagination) (abstract.PaginatedResponse[*authorModel.Author], error) {
	var (
		authorAllDatas     []*authorModel.AuthorData
		authorListResponse abstract.PaginatedResponse[*authorModel.Author]
		totalAuthorsCount  int
		err                error
	)

	err = u.repo.WithTransaction(ctx, func(db database.DataStore) error {
		totalAuthorsCount, err = db.AuthorRepo().CountAuthors(ctx)
		if err != nil {
			return errlist.ParseErrors(err)
		}
		authorListResponse.TotalItems = totalAuthorsCount

		authorAllDatas, err = db.AuthorRepo().ListAuthors(ctx, pgReq.ToPsqlDBStorage())
		if err != nil {
			return errlist.ParseErrors(err)
		}

		return nil
	})
	if err != nil {
		return abstract.PaginatedResponse[*authorModel.Author]{}, nil
	}

	authorList := lo.Map(
		authorAllDatas,
		func(item *authorModel.AuthorData, _ int) *authorModel.Author {
			return item.ToServer()
		},
	)

	authorListResponse.Items = authorList
	authorListResponse.CurrentPage = pgReq.CurrentPage
	authorListResponse.Limit = int(len(authorList))

	return authorListResponse, nil
}
