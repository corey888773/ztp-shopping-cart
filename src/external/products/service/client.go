package service

import (
	"github.com/corey888773/ztp-shopping-cart/src/external/products"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/get_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/v1/remove_from_cart"
)

var _ add_to_cart.ProductsService = (*ProductClientMock)(nil)
var _ remove_from_cart.ProductsService = (*ProductClientMock)(nil)
var _ get_cart.ProductsService = (*ProductClientMock)(nil)

type ProductClientMock struct {
}

func (p ProductClientMock) GetProductByID(productID string) (products.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductClientMock) GetProductsByIDs(productIDs []string) ([]products.Product, error) {
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

func (p ProductClientMock) LockProduct(productID string) error {
	return nil
}

func (p ProductClientMock) UnlockProducts(productIDs []string) error {
	return nil
}
