package usecase

import "github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"

// UseCase struct.
type UseCase struct {
	repo database.DataStore
}

// NewUseCase func.
func NewUseCase(repo database.DataStore) *UseCase {
	return &UseCase{repo: repo}
}
