package get_cart

import (
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

func ApplyEvents(ev []events.CartEvent) *CartBuilder {
	cb := &CartBuilder{}
	for _, event := range ev {
		switch event.EventType {
		case events.EventTypeAddToCart:
			cb.addProduct(event.ProductID, event.Quantity)
		case events.EventTypeRemoveFromCart:
			cb.removeProduct(event.ProductID)
		}
	}
	return cb
}

func (cb *CartBuilder) addProduct(productID string, quantity int) {
	if cb.Products == nil {
		cb.Products = make(map[string]int)
	}
	cb.Products[productID] += quantity
}

func (cb *CartBuilder) removeProduct(productID string) {
	if cb.Products == nil {
		return
	}
	delete(cb.Products, productID)
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
