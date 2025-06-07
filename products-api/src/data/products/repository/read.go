package repository

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/get_all_products"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/get_products"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/unlock_product"
	"gorm.io/gorm"
)

var _ get_products.ReadRepository = (*ReadProductsRepository)(nil)
var _ lock_product.ReadRepository = (*ReadProductsRepository)(nil)
var _ unlock_product.ReadRepository = (*ReadProductsRepository)(nil)
var _ checkout.ReadRepository = (*ReadProductsRepository)(nil)
var _ get_all_products.ReadRepository = (*ReadProductsRepository)(nil)

type ReadProductsRepository struct {
	db *gorm.DB
}

func (r *ReadProductsRepository) GetProductsReservations(productIDs []string) ([]products.ProductReservation, error) {
	var productReservations []products.ProductReservation
	// return max by sequence number for each product
	err := r.db.Table("product_reservations").
		Select("DISTINCT ON (product_id) *").
		Where("product_id IN ?", productIDs).
		Order("product_id, sequence_number desc").
		Find(&productReservations).Error
	if err != nil {
		return nil, err
	}

	return productReservations, nil
}

func (r *ReadProductsRepository) GetProductReservation(productID string) (products.ProductReservation, error) {
	var productReservation products.ProductReservation
	err := r.db.Where("product_id = ?", productID).Order("sequence_number desc").First(&productReservation).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return products.ProductReservation{}, nil // if not found, return empty product reservation
		} else {
			return products.ProductReservation{}, err
		}
	}

	return productReservation, nil
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

// GetAllProducts retrieves all products from the database.
func (r *ReadProductsRepository) GetAllProducts() ([]products.Product, error) {
	var productsList []products.Product
	err := r.db.Find(&productsList).Error
	return productsList, err
}

func (r *ReadProductsRepository) GetProductsByIDs(productIDs []string) ([]products.Product, error) {
	var productsList []products.Product
	err := r.db.Where("id IN ?", productIDs).Find(&productsList).Error
	if err != nil {
		return nil, err
	}

	return productsList, nil
}
