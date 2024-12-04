package manager

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/database"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	authorUC "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author/usecase"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book"
	bookUC "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book/usecase"
)

// DataService interface for all usecases used in this interface.
type DataService interface {
	AuthorService() author.UseCase
	BookService() book.UseCase
}

// Manager Struct.
type ServiceManager struct {
	dataStore database.DataStore
	author    author.UseCase
	book      book.UseCase
}

// NewServiceManager func is.
func NewServiceManager(dataStore database.DataStore) DataService {
	return &ServiceManager{dataStore: dataStore}
}

func (sm *ServiceManager) AuthorService() author.UseCase {
	sm.author = authorUC.NewUseCase(sm.dataStore)
	return sm.author
}

func (sm *ServiceManager) BookService() book.UseCase {
	sm.book = bookUC.NewBookUseCase(sm.dataStore)
	return sm.book
}
