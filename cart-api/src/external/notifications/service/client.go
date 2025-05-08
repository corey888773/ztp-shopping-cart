package notifications

import (
	"fmt"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features/carts/v1/checkout"
)

var _ checkout.NotificationService = (*MockClient)(nil)

type MockClient struct {
}

func (m MockClient) NotifyCheckout(cartID string) error {
	fmt.Println("Mock notification sent for cart ID:", cartID)
	return nil
}
