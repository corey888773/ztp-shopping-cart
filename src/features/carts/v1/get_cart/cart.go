package get_cart

import (
	"encoding/json"

	"github.com/corey888773/ztp-shopping-cart/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/src/external/products"
)

type CartBuilder struct {
	CartID   string         // cartID
	Products map[string]int // productID -> quantity
}

type Product struct {
	ProductDetails products.Product `json:"product_details"`
	Quantity       int              `json:"quantity"`
}

type Cart struct {
	CartID   string    `json:"cart_id"`
	Products []Product `json:"products"`
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
	cb.Products[addToCartPayload.ProductID] += addToCartPayload.Quantity
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
	productsWithDetails := make([]Product, 0, len(cb.Products))
	for _, product := range productDetails {
		if quantity, exists := cb.Products[product.ID]; exists {
			productsWithDetails = append(productsWithDetails, Product{
				ProductDetails: product,
				Quantity:       quantity,
			})
		}
	}

	return &Cart{
		CartID:   cb.CartID,
		Products: productsWithDetails,
	}
}
