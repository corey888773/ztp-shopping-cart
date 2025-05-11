package add_to_cart

type Command struct {
	ProductID string `json:"product_id" binding:"required"`
	CartID    string `json:"cart_id" binding:"required"`
}
