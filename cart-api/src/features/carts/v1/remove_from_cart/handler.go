package remove_from_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
)

type WriteRepository interface {
	RemoveFromCart(cartId string, productId string) error
}

type ReadRepository interface {
	CheckIfCheckedOut(cartID string) (bool, error)
}

type ProductsService interface {
	UnlockProduct(productID string, cartID string) error
}

type Handler struct {
	writeRepository WriteRepository
	productsService ProductsService
	readRepository  ReadRepository
}

func NewHandler(writeRepository WriteRepository, productsSvc ProductsService, readRepository ReadRepository) *Handler {
	return &Handler{
		productsService: productsSvc,
		writeRepository: writeRepository,
		readRepository:  readRepository,
	}
}

func (h *Handler) Handle(command interface{}) error {
	cmd, ok := command.(*Command)
	if !ok {
		return errors.New(commands.ErrInvalidCommand)
	}

	checkedOut, err := h.readRepository.CheckIfCheckedOut(cmd.CartID)
	if err != nil {
		return err
	}

	if checkedOut {
		return errors.New("cart is already checked out")
	}

	err = h.writeRepository.RemoveFromCart(cmd.CartID, cmd.ProductID)
	if err != nil {
		return err
	}

	err = h.productsService.UnlockProduct(cmd.ProductID, cmd.CartID)
	if err != nil {
		return err
	}
	return nil
}
