package checkout

import (
	"errors"
	"time"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/util"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
	"gorm.io/gorm"
)

type UnitOfWork interface {
	Do(fn func(tx *gorm.DB) error) error
}

type ReadRepository interface {
	GetProductsReservations(productIDs []string) ([]products.ProductReservation, error)
}

type WriteRepository interface {
	CheckoutProducts(cartID string, productIdsSequenceNumbersMap map[string]int, tx *gorm.DB) error
}

type Handler struct {
	readRepository  ReadRepository
	writeRepository WriteRepository
	unitOfWork      UnitOfWork
}

func NewHandler(uow UnitOfWork, readRepository ReadRepository, writeRepository WriteRepository) *Handler {
	return &Handler{
		readRepository:  readRepository,
		writeRepository: writeRepository,
		unitOfWork:      uow,
	}
}

func (h *Handler) Handle(command interface{}) error {
	cmd, ok := command.(*Command)
	if !ok {
		return errors.New(ErrInvalidCommand)
	}

	reservations, err := h.readRepository.GetProductsReservations(cmd.ProductIDs)
	if err != nil {
		return err
	}

	if len(reservations) == 0 {
		return errors.New(ErrNoProductsFound)
	}

	if util.Any(reservations, func(res products.ProductReservation) bool {
		return res.CartID != cmd.CartID
	}) {
		return errors.New(ErrProductIsNotLockedToThisCart)
	}

	if util.Any(reservations, func(res products.ProductReservation) bool {
		lockedToTime, err := time.Parse(time.RFC3339, res.LockedToTime)
		if err != nil {
			return false
		}
		return lockedToTime.Before(time.Now())
	}) {
		return errors.New(ErrSomeProductsAreNoLongerLocked)
	}

	productIdsSequenceNumbersMap := make(map[string]int)
	for _, res := range reservations {
		productIdsSequenceNumbersMap[res.ProductID] = res.SequenceNumber
	}

	err = h.unitOfWork.Do(func(tx *gorm.DB) error {
		err := h.writeRepository.CheckoutProducts(cmd.CartID, productIdsSequenceNumbersMap, tx)
		if err != nil {
			return err
		}
		return nil
	})

	return nil
}
