package add_to_cart

const (
	ErrInvalidCommand        = "cannot add to cart: invalid command"
	ErrCartAlreadyCheckedOut = "cannot add to cart: cart is already checked out"
	ErrFailedToLockProduct   = "cannot add to cart: failed to lock product"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrCartAlreadyCheckedOut,
	ErrFailedToLockProduct,
}
