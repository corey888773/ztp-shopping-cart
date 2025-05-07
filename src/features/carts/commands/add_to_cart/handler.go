package add_to_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands"
)

type WriteRepository interface {
	AddToCart(cmd Command) error
}

type ProductsService interface {
	LockProduct(productID string) error
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

	err := h.productsService.LockProduct(cmd.ProductID)
	if err != nil {
		return err
	}

	err = h.repository.AddToCart(*cmd)
	if err != nil {
		_ = h.productsService.UnlockProducts([]string{cmd.ProductID})
		return err
	}

	return nil
}
