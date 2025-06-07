package add_to_cart

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features"
)

func AddToCart(handler commands.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)
	return func(ctx *gin.Context) {
		var addToCartCommand Command
		if err := ctx.ShouldBindJSON(&addToCartCommand); err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		if err := handler.Handle(&addToCartCommand); err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
