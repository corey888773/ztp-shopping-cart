package repository

import (
	"fmt"
	"time"

	"github.com/corey888773/ztp-shopping-cart/src/data/events"
	add_to_cart2 "github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/remove_from_cart"
	"gorm.io/gorm"
)

var _ add_to_cart2.WriteRepository = (*WriteCartRepository)(nil)
var _ remove_from_cart.WriteRepository = (*WriteCartRepository)(nil)

type WriteCartRepository struct {
	db *gorm.DB
}

func NewWriteCartRepository(db *gorm.DB) *WriteCartRepository {
	return &WriteCartRepository{
		db: db,
	}
}

func (w WriteCartRepository) AddToCart(cmd add_to_cart2.Command) error {
	event := events.CartEvent{
		CartID:    cmd.CartID,
		ProductID: cmd.ProductID,
		Quantity:  cmd.Quantity,
	}
	return w.addEventWithRetry(event, events.EventTypeAddToCart, 3)
}

func (w WriteCartRepository) RemoveFromCart(cartID string, productID string) error {
	event := events.CartEvent{
		CartID:    cartID,
		ProductID: productID,
	}
	return w.addEventWithRetry(event, events.EventTypeRemoveFromCart, 3)
}

func (w WriteCartRepository) addEventWithRetry(event events.CartEvent, eventType events.EventType, maxRetries int) error {
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err = w.db.Transaction(func(tx *gorm.DB) error {
			var maxSequence int
			if err := tx.Model(&events.CartEvent{}).
				Where("cart_id = ?", event.CartID).
				Select("COALESCE(MAX(sequence_number), 0)").
				Scan(&maxSequence).Error; err != nil {
				return err
			}

			event.EventType = eventType
			event.SequenceNumber = maxSequence + 1
			event.CreatedAt = time.Now().String()

			return tx.Create(&event).Error
		})

		if err == nil {
			return nil // Success
		}

	}

	return fmt.Errorf("failed to add item to cart after %d attempts: %w", maxRetries, err)
}
