package checkout

type Command struct {
	ProductIDs []string `json:"product_ids"`
	CartID     string   `json:"cart_id"`
}
