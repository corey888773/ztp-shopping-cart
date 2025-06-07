package lock_product

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features"
	"github.com/gin-gonic/gin"
)

func LockProduct(handler commands.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)

	return func(ctx *gin.Context) {
		var cmd Command
		if err := ctx.ShouldBindJSON(&cmd); err != nil {
			errHandler.Handle(ctx, err)
			return
		}
		err := handler.Handle(&cmd)
		if err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
