package commands

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/add_to_cart"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/remove_from_cart"
)

type Handler interface {
	Handle(cmd interface{}) error
}

type WriteRepository interface {
	AddToCart(cmd add_to_cart.Command) error
	RemoveFromCart(cartID string, productID string) error
}

type ProductsService interface {
	LockProduct(productID string) error
	UnlockProducts(productIDs []string) error
}

type CommandHandler struct {
	repository      WriteRepository
	productsService ProductsService
}

func NewCommandHandler(repo WriteRepository, productsSvc ProductsService) *CommandHandler {
	return &CommandHandler{
		productsService: productsSvc,
		repository:      repo,
	}
}

func (h *CommandHandler) Handle(command interface{}) error {
	switch cmd := command.(type) {
	case *add_to_cart.Command:
		return h.HandleAddToCart(cmd)
	case *remove_from_cart.Command:
		return h.HandleRemoveFromCart(cmd)
	default:
		return errors.New(ErrInvalidCommand)
	}
}

func (h *CommandHandler) HandleAddToCart(cmd *add_to_cart.Command) error {
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

func (h *CommandHandler) HandleRemoveFromCart(cmd *remove_from_cart.Command) error {
	return nil
}
