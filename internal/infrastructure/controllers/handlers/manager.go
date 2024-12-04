package handlers

import (
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/controllers"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/infrastructure/manager"
	"github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author"
	authorDelivery "github.com/jumayevgadam/book-app-with-refreshtoken/internal/modules/author/delivery"
)

// DeliveryManager
type DeliveryManager struct {
	author author.Delivery
}

// NewDeliveryManager func.
func NewDeliveryManager(manage manager.DataService) controllers.DataHandlers {
	return &DeliveryManager{
		author: authorDelivery.NewAuthorDelivery(manage),
	}
}

func (dm *DeliveryManager) AuthorDelivery() author.Delivery {
	return dm.author
}
