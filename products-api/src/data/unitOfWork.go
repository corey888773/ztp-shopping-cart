package data

import (
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/lock_product"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features/v1/unlock_product"
	"gorm.io/gorm"
)

var _ lock_product.UnitOfWork = (*UnitOfWork)(nil)
var _ unlock_product.UnitOfWork = (*UnitOfWork)(nil)

type UnitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{
		db: db,
	}
}

func (u *UnitOfWork) Do(fn func(tx *gorm.DB) error) error {
	tx := u.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
