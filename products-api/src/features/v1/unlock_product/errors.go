package unlock_product

const (
	ErrInvalidCommand               = "cannot unlock product: invalid command"
	ErrProductIsNotLockedByThisCart = "cannot unlock product: product is not locked by this cart"
	ErrProductIsNotLocked           = "cannot unlock product: product is not locked"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrProductIsNotLockedByThisCart,
	ErrProductIsNotLocked,
}
