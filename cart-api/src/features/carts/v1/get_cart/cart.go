package get_cart

import (
	"encoding/json"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/external/products"
)

type CartBuilder struct {
	CartID     string         // cartID
	Products   map[string]int // productID -> quantity
	CheckedOut bool           // indicates if cart has been checked out
}


type Cart struct {
	CartID       string             `json:"cart_id"`
	Products     []products.Product `json:"products"`
	IsCheckedOut bool               `json:"is_checked_out"`
}

func NewCartBuilderFromEvents(ev []events.CartEvent) *CartBuilder {
	cb := &CartBuilder{}
	for _, event := range ev {
		switch event.EventType {
		case events.EventTypeAddToCart:
			cb.addProduct(event.Payload)
		case events.EventTypeRemoveFromCart:
			cb.removeProduct(event.Payload)
		case events.EventTypeCheckout:
			// Just stop processing events after checkout
			cb.CheckedOut = true
			return cb
		}
	}
	return cb
}

func (cb *CartBuilder) addProduct(payload string) {
	var addToCartPayload events.AddToCartPayload
	if err := json.Unmarshal([]byte(payload), &addToCartPayload); err != nil {
		return
	}

	if cb.Products == nil {
		cb.Products = make(map[string]int)
	}
	cb.Products[addToCartPayload.ProductID] = 1 // Increment the quantity of the product if it already exists.
	// But for now, we just set it to 1 since we ignore quantity
}

func (cb *CartBuilder) removeProduct(payload string) {
	var removeFromCartPayload events.RemoveFromCartPayload
	if err := json.Unmarshal([]byte(payload), &removeFromCartPayload); err != nil {
		return
	}

	if cb.Products == nil {
		return
	}
	delete(cb.Products, removeFromCartPayload.ProductID)
}

func (cb *CartBuilder) GetProductsList() []string {
	productIDs := make([]string, 0, len(cb.Products))
	for productID := range cb.Products {
		productIDs = append(productIDs, productID)
	}
	return productIDs
}

func (cb *CartBuilder) WithCartID(cartID string) *CartBuilder {
	cb.CartID = cartID
	return cb
}

func (cb *CartBuilder) Build(productDetails []products.Product) *Cart {
	productsWithDetails := make([]products.Product, 0, len(cb.Products))
	for _, product := range productDetails {
		if _, exists := cb.Products[product.ID]; exists {
			productsWithDetails = append(productsWithDetails, product)
		}
	}

	return &Cart{
		CartID:       cb.CartID,
		Products:     productsWithDetails,
		IsCheckedOut: cb.CheckedOut,
	}
}
