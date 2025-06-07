package checkout

const (
	ErrInvalidCommand           = "cannot checkout cart: invalid command"
	ErrCartHasAlreadyCheckedOut = "cannot checkout cart: cart has already checked out"
	ErrFailedToCastToCart       = "cannot checkout cart: failed to cast to cart"
	ErrFailedToCheckoutProducts = "cannot checkout cart: failed to checkout products"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrCartHasAlreadyCheckedOut,
	ErrFailedToCastToCart,
	ErrFailedToCheckoutProducts,
}
