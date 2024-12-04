package repository

import "github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"

// BookRepository struct.
type BookRepository struct {
	psqlDB initializers.DB
}

// NewBookRepository func.
func NewBookRepository(psqlDB initializers.DB) *BookRepository {
	return &BookRepository{psqlDB: psqlDB}
}
