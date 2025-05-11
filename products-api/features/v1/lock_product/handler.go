package lock_product

import (
	"errors"
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/data/products"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Do(fn func(tx *gorm.DB) error) error
}

type WriteRepository interface {
	LockProduct(productID string, tx *gorm.DB) error
}

type ReadRepository interface {
	GetProduct(productID string) (products.Product, error)
}

type Handler struct {
	unitOfWork      UnitOfWork
	writeRepository WriteRepository
	readRepository  ReadRepository
}

func NewHandler(unitOfWork UnitOfWork, writeRepository WriteRepository, readRepository ReadRepository) *Handler {
	return &Handler{
		unitOfWork:      unitOfWork,
		writeRepository: writeRepository,
		readRepository:  readRepository,
	}
}
func (h *Handler) Handle(command interface{}) error {
	cmd, ok := command.(*Command)
	if !ok {
		return errors.New("invalid command")
	}

	err := h.unitOfWork.Do(func(tx *gorm.DB) error {
		prd, err := h.readRepository.GetProduct(cmd.ProductID)
		if err != nil {
			return err
		}

		lockedToTime, err := time.Parse(time.RFC3339, prd.LockedToTime)
		if lockedToTime.After(time.Now()) {
			return errors.New("product is already locked")
		}

		err = h.writeRepository.LockProduct(cmd.ProductID, tx)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
