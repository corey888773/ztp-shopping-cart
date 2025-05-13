package repository

import (
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/checkout"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/features/v1/unlock_product"
	"gorm.io/gorm"
)

var _ lock_product.WriteRepository = (*WriteProductsRepository)(nil)
var _ unlock_product.WriteRepository = (*WriteProductsRepository)(nil)
var _ checkout.WriteRepository = (*WriteProductsRepository)(nil)

type WriteProductsRepository struct {
	db *gorm.DB
}

func (w *WriteProductsRepository) CheckoutProducts(cartID string, productIdsSequenceNumbersMap map[string]int, tx *gorm.DB) error {
	hundredYearsFromNow := time.Now().Add(100 * 365 * 24 * time.Hour).Format(time.RFC3339)
	// create a multiple product reservations with the same cartID and productID. Update their lockedToTime to 100 years from now, and increment the sequence number
	for productID, sequenceNumber := range productIdsSequenceNumbersMap {
		err := tx.Model(&products.ProductReservation{}).Where("product_id = ? AND sequence_number = ?", productID, sequenceNumber).Updates(products.ProductReservation{
			CartID:         cartID,
			LockedToTime:   hundredYearsFromNow,
			SequenceNumber: sequenceNumber + 1,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil

}

func (w *WriteProductsRepository) UnlockProduct(productID string, cartID string, sequenceNumber int, tx *gorm.DB) error {
	lockedToTime := time.Now().Format(time.RFC3339)
	err := tx.Model(&products.ProductReservation{}).Create(&products.ProductReservation{
		ProductID:      productID,
		CartID:         "",
		LockedToTime:   lockedToTime,
		SequenceNumber: sequenceNumber + 1,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (w *WriteProductsRepository) LockProduct(productID string, cartID string, sequenceNumber int, tx *gorm.DB) error {
	lockedToTime := time.Now().Add(15 * time.Minute).Format(time.RFC3339)
	err := tx.Model(&products.ProductReservation{}).Create(&products.ProductReservation{
		ProductID:      productID,
		CartID:         cartID,
		LockedToTime:   lockedToTime,
		SequenceNumber: sequenceNumber + 1,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func NewWriteProductsRepository(db *gorm.DB) *WriteProductsRepository {
	return &WriteProductsRepository{}
}
