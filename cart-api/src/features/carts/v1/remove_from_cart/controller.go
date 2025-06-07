package remove_from_cart

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features"
)

func RemoveFromCart(handler commands.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)
	return func(ctx *gin.Context) {
		var removeFromCart Command
		if err := ctx.ShouldBindJSON(&removeFromCart); err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		if err := handler.Handle(&removeFromCart); err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
