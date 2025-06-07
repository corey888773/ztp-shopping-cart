package checkout

import (
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/features"
)

func Checkout(handler commands.Handler) gin.HandlerFunc {
    errHandler := features.NewErrorHandler(DomainErrors)
    return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := handler.Handle(&Command{CartID: id})
		if err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.Status(204)
	}
}
