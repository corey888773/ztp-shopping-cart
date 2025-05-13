package get_products

import (
	"errors"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
)

type ReadRepository interface {
	GetProductsByIDs(productIDs []string) ([]products.Product, error)
}

type Handler struct {
	readRepository ReadRepository
}

func NewHandler(repo ReadRepository) *Handler {
	return &Handler{
		readRepository: repo,
	}
}

func (h *Handler) Handle(query interface{}) (interface{}, error) {
	qr, ok := query.(*Query)
	if !ok {
		return nil, errors.New(ErrInvalidQuery)
	}

	productsList, err := h.readRepository.GetProductsByIDs(qr.ProductIDs)
	if err != nil {
		return nil, err
	}

	return productsList, nil
}
