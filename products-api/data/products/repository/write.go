package repository

import (
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/unlock_product"
	"gorm.io/gorm"
)

var _ lock_product.WriteRepository = (*WriteProductsRepository)(nil)
var _ unlock_product.WriteRepository = (*WriteProductsRepository)(nil)

type WriteProductsRepository struct {
	db *gorm.DB
}

func (w *WriteProductsRepository) UnlockProduct(productID string, tx *gorm.DB) error {
	lockedToTime := time.Now().Format(time.RFC3339) // Reset the locked_to_time to current time
	err := tx.Model(&products.Product{}).Where("id = ?", productID).Update("locked_to_time", lockedToTime).Error
	if err != nil {
		return err
	}

	return nil
}

func (w *WriteProductsRepository) LockProduct(productID string, tx *gorm.DB) error {
	lockedToTime := time.Now().Add(15 * time.Minute).Format(time.RFC3339)
	err := tx.Model(&products.Product{}).Where("id = ?", productID).Update("locked_to_time", lockedToTime).Error
	if err != nil {
		return err
	}

	return nil
}

func NewWriteProductsRepository(db *gorm.DB) *WriteProductsRepository {
	return &WriteProductsRepository{}
}
