package checkout

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/custom_errors"
	"github.com/gin-gonic/gin"
)

func Checkout(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := handler.Handle(&Command{CartID: id})
		if err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		ctx.Status(204)
	}
}
