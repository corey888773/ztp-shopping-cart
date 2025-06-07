package checkout

const (
	ErrInvalidCommand                = "cannot checkout cart: invalid command"
	ErrProductIsNotLockedToThisCart  = "cannot checkout cart: product is not locked to this cart"
	ErrNoProductsFound               = "cannot checkout cart: no products found"
	ErrSomeProductsAreNoLongerLocked = "cannot checkout cart: some products are no longer locked to this cart"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrProductIsNotLockedToThisCart,
	ErrNoProductsFound,
	ErrSomeProductsAreNoLongerLocked,
}
