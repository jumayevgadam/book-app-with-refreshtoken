package main

import (
	"context"
	"log"

	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/application"
)

func main() {
	err := application.BootStrap(context.Background())
	if err != nil {
		log.Fatal("error in bootstrapping\n")
	}
}
