package get_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/data/events"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/external/products"
)

type ReadRepository interface {
	GetCartEvents(cartID string) ([]events.CartEvent, error)
}

type ProductsService interface {
	GetProductsByIDs(productIDs []string) ([]products.Product, error)
}

type newCartBuilderFunc func([]events.CartEvent) *CartBuilder

type Handler struct {
	repository      ReadRepository
	productsService ProductsService
	newCartBuilder  newCartBuilderFunc
}

func NewHandler(repo ReadRepository, productsSvc ProductsService, newCartBuilderFunc newCartBuilderFunc) *Handler {
	return &Handler{
		productsService: productsSvc,
		repository:      repo,
		newCartBuilder:  newCartBuilderFunc,
	}
}

func (h *Handler) Handle(query interface{}) (interface{}, error) {
	qr, ok := query.(*Query)
	if !ok {
		return nil, errors.New(ErrInvalidQuery)
	}

	evs, err := h.repository.GetCartEvents(qr.CartID)
	if err != nil {
		return nil, err
	}
	cartBuilder := h.newCartBuilder(evs).WithCartID(qr.CartID)

	productsList := cartBuilder.GetProductsList()
	if len(productsList) == 0 {
		return cartBuilder.Build([]products.Product{}), nil
	}

	productDetails, err := h.productsService.GetProductsByIDs(productsList)
	if err != nil {
		return nil, err
	}

	cart := cartBuilder.Build(productDetails)
	return cart, nil
}
