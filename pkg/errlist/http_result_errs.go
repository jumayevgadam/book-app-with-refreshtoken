package errlist

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/config"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/logger"
	"github.com/labstack/echo/v4"
)

// Response returns ErrorResponse, for clean syntax I took function name Response.
// Because in every package i call this errlst package httpError, serviceErr, repoErr.
// Then easily call this httpError.Response(err), serviceErr.Response(err), repoErr.Response(err).
func Response(c echo.Context, err error) error {
	logger := logger.NewApiLogger(&config.Config{})
	logger.InitLogger()

	errStatus, errResponse := ParseErrors(err).Status(), ParseErrors(err)
	logger.Errorf("HTTP Error Response: %v, address: %v", err.Error(), c.Request().RemoteAddr)
	return c.JSON(errStatus, errResponse)
}
