package delivery

import (
	"strconv"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/abstract"
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/manager"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlist"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/reqvalidator"
	"github.com/labstack/echo/v4"
)

var (
	_ author.Delivery = (*AuthorDelivery)(nil)
)

// AuthorDelivery struct.
type AuthorDelivery struct {
	manage manager.DataService
}

// NewAuthorDelivery func.
func NewAuthorDelivery(manage manager.DataService) *AuthorDelivery {
	return &AuthorDelivery{manage: manage}
}

// CreateAuthor delivery.
// @Summary CREATE-AUTHOR.
// @Description create author with properties.
// @Tags AUTHORS
// @ID create-author
// @Accept multipart/form-data
// @Produce json
// @Param request formData authorModel.Request true "author request payload"
// @Success 200 {int} int
// @Failure 400 {object} errlist.RestErr
// @Failure 500 {object} errlist.RestErr
// @Router /author/register [post]
func (d *AuthorDelivery) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request authorModel.Request
		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlist.Response(c, err)
		}

		authorID, err := d.manage.AuthorService().CreateAuthor(c.Request().Context(), request)
		if err != nil {
			return errlist.Response(c, err)
		}

		return c.JSON(200, authorID)
	}
}

// GetAuthor delivery.
// @Summary Get-Author.
// @Description get author by id.
// @Tags Authors
// @ID get-author
// @Accept multipart/form-data
// @Produce json
// @Param author_id path int true "author_id"
// @Success 200 {object} authorModel.Author
// @Failure 400 {object} errlist.RestErr
// @Failure 500 {object} errlist.RestErr
// @Router /author/{author_id} [get]
func (h *AuthorDelivery) GetAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		authorID, err := strconv.Atoi(c.Param("author_id"))
		if err != nil {
			return errlist.Response(c, err)
		}

		author, err := h.manage.AuthorService().GetAuthor(c.Request().Context(), authorID)
		if err != nil {
			return errlist.Response(c, err)
		}

		return c.JSON(200, author)
	}
}

// ListAuthors delivery.
// @Summary List-Authors
// @Description list authors by pagination.
// @Tags Authors
// @ID list-authors
// @Accept multipart/form-data
// @Produce json
// @Param current-page query int false "current-page number" Format(current-page)
// @Param limit query int false "number of items per page" Format(limit)
// @Success 200 {object} []authorModel.Auhtor
// @Failure 400 {object} errlist.RestErr
// @Failure 500 {object} errlist.RestErr
// @Router /author/list [get]
func (h *AuthorDelivery) ListAuthors() echo.HandlerFunc {
	return func(c echo.Context) error {
		pgReq, err := abstract.GetPaginationFromContext(c)
		if err != nil {
			return errlist.Response(c, err)
		}

		authorList, err := h.manage.AuthorService().ListAuthors(c.Request().Context(), pgReq)
		if err != nil {
			return errlist.Response(c, err)
		}

		return c.JSON(200, authorList)
	}
}
