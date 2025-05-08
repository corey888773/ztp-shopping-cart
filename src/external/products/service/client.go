package products

import (
	"github.com/corey888773/ztp-shopping-cart/src/external/products"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/remove_from_cart"
)

var _ add_to_cart.ProductsService = (*ClientMock)(nil)
var _ remove_from_cart.ProductsService = (*ClientMock)(nil)
var _ get_cart.ProductsService = (*ClientMock)(nil)
var _ checkout.ProductsService = (*ClientMock)(nil)

type ClientMock struct {
}

func (p ClientMock) CheckoutProducts(productIDs []string) error {
	return nil
}

func (p ClientMock) GetProductByID(productID string) (products.Product, error) {
	//TODO implement me
	panic("implement me")
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

func (p ClientMock) LockProduct(productID string) error {
	return nil
}

func (p ClientMock) UnlockProducts(productIDs []string) error {
	return nil
}
