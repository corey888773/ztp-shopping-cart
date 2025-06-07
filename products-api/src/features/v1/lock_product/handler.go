package lock_product

import (
	"errors"
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/common/util"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Do(fn func(tx *gorm.DB) error) error
}

type WriteRepository interface {
	LockProduct(productID string, cartID string, seqNumber int, tx *gorm.DB) error
}

type ReadRepository interface {
	GetProductReservation(productID string) (products.ProductReservation, error)
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
		return errors.New(ErrInvalidCommand)
	}

	action := func() error {
		return h.unitOfWork.Do(func(tx *gorm.DB) error {
			_, err := h.readRepository.GetProduct(cmd.ProductID)
			if err != nil {
				return err
			}

			prdRes, err := h.readRepository.GetProductReservation(cmd.ProductID)
			if err != nil {
				return err
			}

			lockedToTime, err := time.Parse(time.RFC3339, prdRes.LockedToTime)
			if err != nil {
				return err
			}

			if lockedToTime.After(time.Now()) {
				return errors.New(ErrProductIsAlreadyLocked)
			}

			err = h.writeRepository.LockProduct(cmd.ProductID, cmd.CartID, prdRes.SequenceNumber, tx)
			if err != nil {
				return err
			}
			return nil
		})
	}

	if err := util.InvokeWithRetry(action, 1*time.Second, 3); err != nil {
		return err
	}

	return nil
}
