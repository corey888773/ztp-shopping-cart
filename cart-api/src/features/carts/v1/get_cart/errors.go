package get_cart

const (
	ErrInvalidQuery    = "cannot get cart: invalid query"
	ErrNoProductsFound = "cannot get cart: no products found"
)

var DomainErrors = []string{
	ErrInvalidQuery,
	ErrNoProductsFound,
}
