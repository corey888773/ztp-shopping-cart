package repository

import (
	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/get_products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/unlock_product"
	"gorm.io/gorm"
)

var _ get_products.ReadRepository = (*ReadProductsRepository)(nil)
var _ lock_product.ReadRepository = (*ReadProductsRepository)(nil)
var _ unlock_product.ReadRepository = (*ReadProductsRepository)(nil)

type ReadProductsRepository struct {
	db *gorm.DB
}

func (r *ReadProductsRepository) GetProduct(productID string) (products.Product, error) {
	var product products.Product
	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return products.Product{}, err
	}

	return product, nil
}

func NewReadProductsRepository(db *gorm.DB) *ReadProductsRepository {
	return &ReadProductsRepository{
		db: db,
	}
}

func (r *ReadProductsRepository) GetProductsByIDs(productIDs []string) ([]products.Product, error) {
	var productsList []products.Product
	err := r.db.Where("id IN ?", productIDs).Find(&productsList).Error
	if err != nil {
		return nil, err
	}

	return productsList, nil
}
