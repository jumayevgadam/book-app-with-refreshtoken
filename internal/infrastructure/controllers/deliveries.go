package controllers

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
)

// Handlers interface for all delivery layers of app.
type DataHandlers interface {
	AuthorDelivery() author.Delivery
}
