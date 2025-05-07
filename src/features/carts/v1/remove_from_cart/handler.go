package remove_from_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/src/common/commands"
)

type WriteRepository interface {
	RemoveFromCart(cartId string, productId string) error
}

type ProductsService interface {
	UnlockProducts(productIDs []string) error
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

	err := h.repository.RemoveFromCart(cmd.CartID, cmd.ProductID)
	if err != nil {
		return err
	}

	err = h.productsService.UnlockProducts([]string{cmd.ProductID})
	if err != nil {
		return err
	}
	return nil
}
