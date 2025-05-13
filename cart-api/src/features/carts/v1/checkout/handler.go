package checkout

import (
	"errors"
	"fmt"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/util"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/get_cart"
)

type WriteRepository interface {
	Checkout(command Command) error
}

type ReadRepository interface {
	GetCartEvents(cartID string) ([]events.CartEvent, error)
}

type NotificationService interface {
	NotifyCheckout(cartID string) error
}

type ProductsService interface {
	CheckoutProducts(productIDs []string, cartID string) error
}

type GetCartQueryHandler interface {
	Handle(query interface{}) (interface{}, error)
}

type Handler struct {
	writeRepository     WriteRepository
	readRepository      ReadRepository
	productsService     ProductsService
	getCartQuery        GetCartQueryHandler
	notificationService NotificationService
}

func NewHandler(
	writeRepository WriteRepository,
	productsSvc ProductsService,
	getCartQuery GetCartQueryHandler,
	notificationService NotificationService,
	readRepository ReadRepository) *Handler {
	return &Handler{
		writeRepository:     writeRepository,
		readRepository:      readRepository,
		productsService:     productsSvc,
		getCartQuery:        getCartQuery,
		notificationService: notificationService,
	}
}

func (h *Handler) Handle(command interface{}) error {
	cmd, ok := command.(*Command)
	if !ok {
		return errors.New(commands.ErrInvalidCommand)
	}

	evs, err := h.readRepository.GetCartEvents(cmd.CartID)
	fmt.Printf("events: %+v \n", evs)
	if err != nil {
		return err
	}

	if alreadyCheckedOut(evs) {
		return errors.New("cart has already been checked out")
	}

	err = h.writeRepository.Checkout(*cmd)
	if err != nil {
		return err
	}

	err = h.notificationService.NotifyCheckout(cmd.CartID)
	if err != nil {
		return err
	}

	cart, err := h.getCartQuery.Handle(&get_cart.Query{CartID: cmd.CartID})
	if err != nil {
		return err
	}

	cartDetails, ok := cart.(*get_cart.Cart)
	if !ok {
		return errors.New("failed to cast cart to Cart type")
	}

	productIDs := util.Map(cartDetails.Products, func(p get_cart.Product) string {
		return p.ProductDetails.ID
	})

	err = h.productsService.CheckoutProducts(productIDs, cmd.CartID)
	if err != nil {
		return err
	}

	return nil
}

func alreadyCheckedOut(evs []events.CartEvent) bool {
	return util.Any(evs, func(ev events.CartEvent) bool {
		return ev.EventType == events.EventTypeCheckout
	})
}
