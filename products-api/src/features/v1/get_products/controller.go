package get_products

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features"
	"github.com/gin-gonic/gin"
)

func GetProducts(handler queries.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)

	return func(ctx *gin.Context) {
		var getProductsQuery Query
		if err := ctx.ShouldBindJSON(&getProductsQuery); err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		products, err := handler.Handle(&getProductsQuery)
		if err != nil {
			errHandler.Handle(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}
