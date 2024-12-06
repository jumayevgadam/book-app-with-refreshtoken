package delivery

import (
	authorModel "github.com/jumayevgadam/book-app-with-refreshtoken/internal/domain/models/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/manager"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
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
// @Failure 400 {object} errlst.RestErr
// @Failure 500 {object} errlst.RestErr
// @Router /author/register [post]
func (d *AuthorDelivery) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request authorModel.Request
		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return errlst.Response(c, err)
		}

		authorID, err := d.manage.AuthorService().CreateAuthor(c.Request().Context(), request)
		if err != nil {
			return errlst.Response(c, err)
		}

		return c.JSON(200, authorID)
	}
}

// GetAuthor delivery.