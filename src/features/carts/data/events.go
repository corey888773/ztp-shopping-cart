package data

type EventType string

const (
	EventTypeAddToCart      EventType = "add_to_cart"
	EventTypeRemoveFromCart EventType = "remove_from_cart"
	EventTypeUpdateQuantity EventType = "update_quantity"
	EventTypeCheckout       EventType = "checkout"
)

type CartEvent struct {
	CartID         string    `json:"cart_id" gorm:"primaryKey;column:cart_id"`
	SequenceNumber int       `json:"sequence_number" gorm:"primaryKey;column:sequence_number;autoIncrement:false"`
	ProductID      string    `json:"product_id" gorm:"column:product_id"`
	EventType      EventType `json:"event_type" gorm:"column:event_type"`
	Quantity       int       `json:"quantity" gorm:"column:quantity"`
	CreatedAt      string    `json:"created_at" gorm:"column:created_at"`
}
