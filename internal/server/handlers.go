package server

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/docs"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/controllers/handlers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/manager/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// MapHandlers.
func (s *Server) MapHandlers() error {
	docs.SwaggerInfo.Title = "API DOCUMENTATION OF BOOK-APP-WITH-REFRESH-TOKEN"
	s.Echo.GET("/api-docs/swagger/book-app-with-refreshToken/*", echoSwagger.WrapHandler)
	v1 := s.Echo.Group("/api/v1")

	UseCases := service.NewServiceManager(s.DataStore)

	Handlers := handlers.NewDeliveryManager(UseCases)

	// for authors.
	authorGroup := v1.Group("/author")
	{
		authorGroup.POST("/register", Handlers.AuthorDelivery().CreateAuthor())
	}

	// for books.

	return nil
}
