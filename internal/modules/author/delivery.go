package author

import "github.com/labstack/echo/v4"

// Delivery interface for authors.
type Delivery interface {
	CreateAuthor() echo.HandlerFunc
	GetAuthor() echo.HandlerFunc
	ListAuthors() echo.HandlerFunc
}
