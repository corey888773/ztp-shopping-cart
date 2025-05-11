package repository

import (
	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/get_products"
	"gorm.io/gorm"
)

var _ get_products.ReadRepository = (*ReadProductsRepository)(nil)

type ReadProductsRepository struct {
	db *gorm.DB
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
