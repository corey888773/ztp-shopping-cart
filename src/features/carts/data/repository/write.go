package repository

import (
	"fmt"
	"time"

	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/remove_from_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/data"
	"gorm.io/gorm"
)

var _ add_to_cart.WriteRepository = (*WriteCartRepository)(nil)
var _ remove_from_cart.WriteRepository = (*WriteCartRepository)(nil)

type WriteCartRepository struct {
	db *gorm.DB
}

func NewWriteCartRepository(db *gorm.DB) *WriteCartRepository {
	return &WriteCartRepository{
		db: db,
	}
}

func (w WriteCartRepository) AddToCart(cmd add_to_cart.Command) error {
	event := data.CartEvent{
		CartID:    cmd.CartID,
		ProductID: cmd.ProductID,
		Quantity:  cmd.Quantity,
	}
	return w.addEventWithRetry(event, data.EventTypeAddToCart, 3)
}

func (w WriteCartRepository) RemoveFromCart(cartID string, productID string) error {
	event := data.CartEvent{
		CartID:    cartID,
		ProductID: productID,
	}
	return w.addEventWithRetry(event, data.EventTypeRemoveFromCart, 3)
}

func (w WriteCartRepository) addEventWithRetry(event data.CartEvent, eventType data.EventType, maxRetries int) error {
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err = w.db.Transaction(func(tx *gorm.DB) error {
			var maxSequence int
			if err := tx.Model(&data.CartEvent{}).
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
