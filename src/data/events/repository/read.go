package repository

import (
	"github.com/corey888773/ztp-shopping-cart/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/get_cart"
	"gorm.io/gorm"
)

var _ get_cart.ReadRepository = (*ReadCartRepository)(nil)

type ReadCartRepository struct {
	db *gorm.DB
}

func NewReadCartRepository(db *gorm.DB) *ReadCartRepository {
	return &ReadCartRepository{
		db: db,
	}
}

func (r ReadCartRepository) GetCartEvents(cartID string) ([]events.CartEvent, error) {
	var events []events.CartEvent
	err := r.db.Where("cart_id = ?", cartID).Find(&events).Order("sequence_number").Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
