package checkout

type Command struct {
	CartID string `json:"cart_id" binding:"required"`
}
