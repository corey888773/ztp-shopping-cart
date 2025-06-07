package lock_product

const (
	ErrInvalidCommand         = "cannot lock product: invalid command"
	ErrProductIsAlreadyLocked = "cannot lock product: product is already locked"
)

var DomainErrors = []string{
	ErrInvalidCommand,
	ErrProductIsAlreadyLocked,
}
