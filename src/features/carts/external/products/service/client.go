package service

import (
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/external/products"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries"
)

var _ commands.ProductsService = (*ProductClientMock)(nil)
var _ queries.ProductsService = (*ProductClientMock)(nil)

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
