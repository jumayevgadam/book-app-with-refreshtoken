package manager

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/book"
)

// DataService interface for all usecases used in this interface.
type DataService interface {
	AuthorService() author.UseCase
	BookService() book.UseCase
}
