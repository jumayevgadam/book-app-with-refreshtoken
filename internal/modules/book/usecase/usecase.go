package usecase

import "github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"

// BookUseCase struct.
type BookUseCase struct {
	repo database.DataStore
}

// NewBookUseCase func.
func NewBookUseCase(repo database.DataStore) *BookUseCase {
	return &BookUseCase{repo: repo}
}
