package get_cart

type Query struct {
	CartID string `json:"cart_id" binding:"required"`
}
