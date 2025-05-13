package products

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/external/products"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/remove_from_cart"
)

var _ add_to_cart.ProductsService = (*ClientMock)(nil)
var _ remove_from_cart.ProductsService = (*ClientMock)(nil)
var _ get_cart.ProductsService = (*ClientMock)(nil)
var _ checkout.ProductsService = (*ClientMock)(nil)

type ClientMock struct {
}

func (p ClientMock) CheckoutProducts(productIDs []string, cartID string) error {
	return nil
}

func (p ClientMock) GetProductsByIDs(productIDs []string) ([]products.Product, error) {
	return []products.Product{
		{
			ID:   "1",
			Name: "Product 1",
		},
		{
			ID:   "2",
			Name: "Product 2",
		},
		{
			ID:   "3",
			Name: "Product 3",
		},
	}, nil
}

func (p ClientMock) LockProduct(productID string, cartID string) error {
	return nil
}

func (p ClientMock) UnlockProduct(productID string, cartID string) error {
	return nil
}
