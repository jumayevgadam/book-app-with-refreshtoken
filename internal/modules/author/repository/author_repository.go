package repository

import "github.com/jumayevgadam/book-app-with-refreshtoken/internal/initializers"

// AuthorRepository struct.
type AuthorRepository struct {
	psqlDB initializers.DB
}

// NewAuthorRepository func.
func NewAuthorRepository(psqlDB initializers.DB) *AuthorRepository {
	return &AuthorRepository{psqlDB: psqlDB}
}
