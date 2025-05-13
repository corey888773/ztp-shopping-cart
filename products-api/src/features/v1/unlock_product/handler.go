package unlock_product

import (
	"errors"
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Do(fn func(tx *gorm.DB) error) error
}

type WriteRepository interface {
	UnlockProduct(productID string, cartID string, sequenceNumber int, tx *gorm.DB) error
}

type ReadRepository interface {
	GetProduct(productID string) (products.Product, error)
	GetProductReservation(productID string) (products.ProductReservation, error)
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
		return errors.New(ErrInvalidCommand)
	}

	err := h.unitOfWork.Do(func(tx *gorm.DB) error {
		_, err := h.readRepository.GetProduct(cmd.ProductID)
		if err != nil {
			return err
		}

		prdRes, err := h.readRepository.GetProductReservation(cmd.ProductID)
		if err != nil {
			return err
		}

		if prdRes.CartID != cmd.CartID {
			return errors.New(ErrProductIsNotLockedByThisCart)
		}

		lockedToTime, err := time.Parse(time.RFC3339, prdRes.LockedToTime)
		if lockedToTime.Before(time.Now()) {
			return errors.New(ErrProductIsNotLocked)
		}

		err = h.writeRepository.UnlockProduct(cmd.ProductID, cmd.CartID, prdRes.SequenceNumber, tx)
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
