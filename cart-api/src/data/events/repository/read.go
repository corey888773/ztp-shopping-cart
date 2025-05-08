package repository

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/get_cart"
	"gorm.io/gorm"
)

var _ get_cart.ReadRepository = (*ReadCartRepository)(nil)
var _ checkout.ReadRepository = (*ReadCartRepository)(nil)

type ReadCartRepository struct {
	db *gorm.DB
}

func NewReadCartRepository(db *gorm.DB) *ReadCartRepository {
	return &ReadCartRepository{
		db: db,
	}
}

func (r ReadCartRepository) GetCartEvents(cartID string) ([]events.CartEvent, error) {
	var evs []events.CartEvent
	err := r.db.Where("cart_id = ?", cartID).Find(&evs).Order("sequence_number").Error
	if err != nil {
		return nil, err
	}
	return evs, nil
}
