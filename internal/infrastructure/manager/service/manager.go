package service

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/manager"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	authorUC "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author/usecase"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book"
	bookUC "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book/usecase"
)

// Manager Struct.
type ServiceManager struct {
	author author.UseCase
	book   book.UseCase
}

// NewServiceManager func is.
func NewServiceManager(dataStore database.DataStore) manager.DataService {
	return &ServiceManager{
		author: authorUC.NewAuthorUseCase(dataStore),
		book:   bookUC.NewBookUseCase(dataStore),
	}
}

// AuthorService returns the author use case.
func (sm *ServiceManager) AuthorService() author.UseCase {
	return sm.author
}

// BookService returns the book use case.
func (sm *ServiceManager) BookService() book.UseCase {
	return sm.book
}
