package remove_from_cart

const (
	ErrInvalidCommand        = "cannot remove from cart: invalid command"
	ErrCartAlreadyCheckedOut = "cannot remove from cart: cart already checked out"
	ErrFailedToUnlockProduct = "cannot remove from cart: failed to unlock product"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrCartAlreadyCheckedOut,
	ErrFailedToUnlockProduct,
}
