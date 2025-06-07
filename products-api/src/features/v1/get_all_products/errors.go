package get_all_products

// Error constants for get_all_products feature.
const (
	ErrInvalidQuery = "cannot fetch all products: invalid query"
)

var DomainErrors = []string{
	ErrInvalidQuery,
}
