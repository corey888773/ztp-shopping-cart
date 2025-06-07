package get_all_products

import (
	"errors"
	"time"

	"github.com/corey888773/ztp-shopping-cart/products-api/src/data/products"
)

// ReadRepository defines the interface for reading products and their reservations.
type ReadRepository interface {
	GetAllProducts() ([]products.Product, error)
	GetProductsReservations(productIDs []string) ([]products.ProductReservation, error)
}

// Handler handles the get_all_products query.
type Handler struct {
	readRepository ReadRepository
}

// ProductWithAvailability wraps a product with its availability status.
type ProductWithAvailability struct {
	ID          string `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAvailable bool   `json:"is_available"`
}

// NewHandler constructs a new Handler.
func NewHandler(repo ReadRepository) *Handler {
	return &Handler{
		readRepository: repo,
	}
}

// Handle executes the query to get all products with availability status.
func (h *Handler) Handle(query interface{}) (interface{}, error) {
	_, ok := query.(*Query)
	if !ok {
		return nil, errors.New(ErrInvalidQuery)
	}

	// Fetch all products.
	productsList, err := h.readRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	// Collect product IDs.
	productIDs := make([]string, len(productsList))
	for i, p := range productsList {
		productIDs[i] = p.ID
	}

	// Fetch reservations for products.
	reservations, err := h.readRepository.GetProductsReservations(productIDs)
	if err != nil {
		return nil, err
	}

	// Map reservations by product ID.
	resMap := make(map[string]products.ProductReservation)
	for _, res := range reservations {
		resMap[res.ProductID] = res
	}

	// Determine availability.
	now := time.Now().UTC()
	var result []ProductWithAvailability
	for _, p := range productsList {
		available := true
		if res, exists := resMap[p.ID]; exists {
			if lockedTime, err := time.Parse(time.RFC3339, res.LockedToTime); err == nil && lockedTime.After(now) {
				available = false
			}
		}
		result = append(result, ProductWithAvailability{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			IsAvailable: available,
		})
	}

	return result, nil
}
