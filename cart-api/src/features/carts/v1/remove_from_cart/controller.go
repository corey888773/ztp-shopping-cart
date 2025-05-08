package remove_from_cart

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/custom_errors"
	"github.com/gin-gonic/gin"
)

func RemoveFromCart(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var removeFromCart Command
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
