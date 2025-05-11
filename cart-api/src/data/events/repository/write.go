package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/remove_from_cart"
	"gorm.io/gorm"
)

var _ add_to_cart.WriteRepository = (*WriteCartRepository)(nil)
var _ remove_from_cart.WriteRepository = (*WriteCartRepository)(nil)
var _ checkout.WriteRepository = (*WriteCartRepository)(nil)

type WriteCartRepository struct {
	db *gorm.DB
}

func (w WriteCartRepository) Checkout(command checkout.Command) error {
	event := events.CartEvent{
		CartID:    command.CartID,
		EventType: events.EventTypeCheckout,
	}
	return w.addEventWithRetry(event, 3)
}

func NewWriteCartRepository(db *gorm.DB) *WriteCartRepository {
	return &WriteCartRepository{
		db: db,
	}
}

func (w WriteCartRepository) AddToCart(cmd add_to_cart.Command) error {
	payload, err := json.Marshal(events.AddToCartPayload{
		ProductID: cmd.ProductID,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	event := events.CartEvent{
		CartID:    cmd.CartID,
		EventType: events.EventTypeAddToCart,
		Payload:   string(payload),
	}
	return w.addEventWithRetry(event, 3)
}

func (w WriteCartRepository) RemoveFromCart(cartID string, productID string) error {
	payload, err := json.Marshal(events.RemoveFromCartPayload{
		ProductID: productID,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	event := events.CartEvent{
		CartID:    cartID,
		EventType: events.EventTypeRemoveFromCart,
		Payload:   string(payload),
	}
	return w.addEventWithRetry(event, 3)
}

func (w WriteCartRepository) addEventWithRetry(event events.CartEvent, maxRetries int) error {
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
