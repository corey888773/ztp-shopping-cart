package repository

import (
	"errors"
	"fmt"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/remove_from_cart"
	"gorm.io/gorm"
)

var _ get_cart.ReadRepository = (*ReadCartRepository)(nil)
var _ checkout.ReadRepository = (*ReadCartRepository)(nil)
var _ add_to_cart.ReadRepository = (*ReadCartRepository)(nil)
var _ remove_from_cart.ReadRepository = (*ReadCartRepository)(nil)

type ReadCartRepository struct {
	db *gorm.DB
}

func (r ReadCartRepository) CheckIfCheckedOut(cartID string) (bool, error) {
	var lastEvent events.CartEvent
	err := r.db.Where("cart_id = ?", cartID).Order("sequence_number desc").First(&lastEvent).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	fmt.Printf("lastEvent: %+v\n", lastEvent)
	if lastEvent.EventType == events.EventTypeCheckout {
		return true, nil
	}

	return false, nil
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
