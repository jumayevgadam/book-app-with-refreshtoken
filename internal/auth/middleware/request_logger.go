package middleware

import (
	"time"

	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/errlst"
	"github.com/jumayevgadam/book-app-with-refreshtoken/pkg/utils"
	"github.com/labstack/echo/v4"
)

// RequestLoggerMiddleware for logging http requests.
func (mw *MiddlewareManager) RequestLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)

		req := c.Request()
		res := c.Response()
		status := res.Status
		size := res.Size

		s := time.Since(start).String()
		requestID := utils.GetRequestID(c)

		mw.logger.Infof("RequestID: %s, Method: %s, URI: %s, Status: %v, Size: %v, Time: %s",
			requestID, req.Method, req.URL, status, size, s)

		return errlst.Response(c, err)
	}
}
