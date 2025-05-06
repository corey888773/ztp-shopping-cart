package queries

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/src/features/carts/data"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/external/products"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries/get_cart"
)

type ReadRepository interface {
	GetCartEvents(cartID string) ([]data.CartEvent, error)
}

type ProductsService interface {
	GetProductByID(productID string) (products.Product, error)
	GetProductsByIDs(productIDs []string) ([]products.Product, error)
}

type newCartBuilderFunc func([]data.CartEvent) *get_cart.CartBuilder

type Handler interface {
	Handle(query interface{}) (interface{}, error)
}

type QueryHandler struct {
	repository      ReadRepository
	productsService ProductsService
	newCartBuilder  newCartBuilderFunc
}

func NewQueryHandler(repo ReadRepository, productsSvc ProductsService, newCartBuilderFunc newCartBuilderFunc) *QueryHandler {
	return &QueryHandler{
		productsService: productsSvc,
		repository:      repo,
		newCartBuilder:  newCartBuilderFunc,
	}
}

func (h *QueryHandler) Handle(query interface{}) (interface{}, error) {
	switch q := query.(type) {
	case *get_cart.Query:
		return h.handleGetCart(q)
	default:
		return nil, errors.New(ErrInvalidQuery)
	}
}

func (h *QueryHandler) handleGetCart(query *get_cart.Query) (interface{}, error) {
	events, err := h.repository.GetCartEvents(query.CartID)
	if err != nil {
		return nil, err
	}
	cartBuilder := h.newCartBuilder(events)

	productDetails, err := h.productsService.GetProductsByIDs(cartBuilder.GetProductsList())
	if err != nil {
		return nil, err
	}

	cart := cartBuilder.WithCartID(query.CartID).Build(productDetails)
	return cart, nil
}
