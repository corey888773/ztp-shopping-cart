package get_cart

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/src/features/carts/data"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/external/products"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries"
)

type ReadRepository interface {
	GetCartEvents(cartID string) ([]data.CartEvent, error)
}

type ProductsService interface {
	GetProductByID(productID string) (products.Product, error)
	GetProductsByIDs(productIDs []string) ([]products.Product, error)
}

type newCartBuilderFunc func([]data.CartEvent) *CartBuilder

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
		return nil, errors.New(queries.ErrInvalidQuery)
	}

	events, err := h.repository.GetCartEvents(qr.CartID)
	if err != nil {
		return nil, err
	}
	cartBuilder := h.newCartBuilder(events)

	productDetails, err := h.productsService.GetProductsByIDs(cartBuilder.GetProductsList())
	if err != nil {
		return nil, err
	}

	cart := cartBuilder.WithCartID(qr.CartID).Build(productDetails)
	return cart, nil
}
