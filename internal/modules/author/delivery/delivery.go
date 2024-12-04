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
func (d *AuthorDelivery) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request authorModel.Request
		err := reqvalidator.ReadRequest(c, &request)
		if err != nil {
			return c.JSON(400, errlst.ErrBadRequest)
		}

		authorID, err := d.manage.AuthorService().CreateAuthor(c.Request().Context(), request)
		if err != nil {
			return c.JSON(500, errlst.ErrInternalServer)
		}

		return c.JSON(200, authorID)
	}
}
