package repository

import (
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/data"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries"
	"gorm.io/gorm"
)

var _ queries.ReadRepository = (*ReadCartRepository)(nil)

type ReadCartRepository struct {
	db *gorm.DB
}

func NewReadCartRepository(db *gorm.DB) *ReadCartRepository {
	return &ReadCartRepository{
		db: db,
	}
}

func (r ReadCartRepository) GetCartEvents(cartID string) ([]data.CartEvent, error) {
	var events []data.CartEvent
	err := r.db.Where("cart_id = ?", cartID).Find(&events).Order("sequence_number").Error
	if err != nil {
		return nil, err
	}
	return events, nil
}
