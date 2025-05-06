package get_cart

import (
	"github.com/corey888773/ztp-shopping-cart/src/common/custom_errors"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries"
	"github.com/corey888773/ztp-shopping-cart/src/features/carts/queries/get_cart"
	"github.com/gin-gonic/gin"
)

func GetCart(handler queries.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, err := handler.Handle(&get_cart.Query{CartID: id})
		if err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		ctx.JSON(200, result)
	}
}
