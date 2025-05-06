package add_to_cart

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/src/common/custom_errors"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/commands/add_to_cart"
	"github.com/gin-gonic/gin"
)

func AddToCart(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var addToCartCommand add_to_cart.Command
		if err := ctx.ShouldBindJSON(&addToCartCommand); err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		if err := handler.Handle(&addToCartCommand); err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
