package events

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
	EventType      EventType `json:"event_type" gorm:"column:event_type"`
	CreatedAt      string    `json:"created_at" gorm:"column:created_at"`
	Payload        string    `json:"payload" gorm:"column:payload"`
}

type RemoveFromCartPayload struct {
	ProductID string `json:"product_id"`
}

type AddToCartPayload struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
