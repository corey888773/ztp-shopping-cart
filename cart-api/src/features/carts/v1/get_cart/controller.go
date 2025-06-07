package get_cart

import (
    "github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
    "github.com/gin-gonic/gin"
    "github.com/corey888773/ztp-shopping-cart/cart-api/src/features"
)

func GetCart(handler queries.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		result, err := handler.Handle(&Query{CartID: id})
		if err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.JSON(200, result)
	}
}
