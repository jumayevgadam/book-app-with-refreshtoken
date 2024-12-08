package abstract

import (
	"strconv"

	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
	"github.com/labstack/echo/v4"
)

// PaginationQuery is.
type Pagination struct {
	Limit       int `form:"limit" json:"limit" validate:"required,gte=0,lte=100"`
	CurrentPage int `form:"current-page" json:"currentPage" validate:"required"`
}

// PaginationData
type PaginationData struct {
	Limit       int `db:"limit"`
	CurrentPage int `db:"current_page"`
}

type PaginatedResponse[T any] struct {
	Items       []T `json:"items"`
	Limit       int `json:"limit"`
	CurrentPage int `json:"currentPage"`
	TotalItems  int `json:"totalItems"`
}

// ToStorage sends pagination psqlDBStorage.
func (p *Pagination) ToPsqlDBStorage() PaginationData {
	return PaginationData{
		Limit:       p.Limit,
		CurrentPage: p.CurrentPage,
	}
}

func (pg *Pagination) SetLimit(limit string) error {
	if limit == "" {
		pg.Limit = 10
		return nil
	}

	num, err := strconv.Atoi(limit)
	if err != nil {
		return errlist.ParseErrors(err)
	}
	pg.Limit = num

	return nil
}

func (pg *Pagination) SetCurrentPage(currentPage string) error {
	if currentPage == "" {
		pg.CurrentPage = 1
		return nil
	}

	num, err := strconv.Atoi(currentPage)
	if err != nil {
		return errlist.ParseErrors(err)
	}
	pg.CurrentPage = num

	return nil
}

// GetPaginationFromContext.
func GetPaginationFromContext(c echo.Context) (Pagination, error) {
	pg := Pagination{}

	if err := pg.SetCurrentPage(c.QueryParam("current-page")); err != nil {
		return pg, errlist.ParseErrors(err)
	}

	if err := pg.SetLimit(c.QueryParam("limit")); err != nil {
		return pg, errlist.ParseErrors(err)
	}

	return pg, nil
}
