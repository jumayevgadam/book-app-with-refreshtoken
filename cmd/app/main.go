package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/application"
)

// @title BOOK-APP-WITH-REFRESH-TOKEN api documentation
// @version 2.0
// @description book app with refresh token.
// @termsOfService http://swagger.io/terms
// @contact.name Gadam Jumayev
// @contact.url https://github.com/jumayevgadam
// @contact.email hypergadam@gmail.com
// @host localhost:4000
// @BasePath /api/v1
func main() {
	err := application.BootStrap(context.Background())
	if err != nil {
		log.Fatal("error in bootstrapping\n")
	}
}
