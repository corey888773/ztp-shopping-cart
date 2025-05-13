package add_to_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
)

type WriteRepository interface {
	AddToCart(cmd Command) error
}

type ProductsService interface {
	LockProduct(productID string, cartID string) error
	UnlockProduct(productID string, cartID string) error
}

type Handler struct {
	repository      WriteRepository
	productsService ProductsService
}

func NewHandler(repo WriteRepository, productsSvc ProductsService) *Handler {
	return &Handler{
		productsService: productsSvc,
		repository:      repo,
	}
}

func (h *Handler) Handle(command interface{}) error {
	cmd, ok := command.(*Command)
	if !ok {
		return errors.New(commands.ErrInvalidCommand)
	}

	err := h.productsService.LockProduct(cmd.ProductID, cmd.CartID)
	if err != nil {
		return err
	}

	err = h.repository.AddToCart(*cmd)
	if err != nil {
		_ = h.productsService.UnlockProduct(cmd.ProductID, cmd.CartID)
		return err
	}

	return nil
}
