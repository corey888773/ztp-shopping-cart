package remove_from_cart

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/src/common/custom_errors"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/remove_from_cart"
	"github.com/gin-gonic/gin"
)

func RemoveFromCart(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var removeFromCart remove_from_cart.Command
		if err := ctx.ShouldBindJSON(&removeFromCart); err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		if err := handler.Handle(&removeFromCart); err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
