package get_products

import (
	"fmt"
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/custom_errors"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
	"github.com/gin-gonic/gin"
)

func GetProducts(handler queries.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var getProductsQuery Query
		fmt.Printf("getProductsQuery: %+v\n", getProductsQuery)
		if err := ctx.ShouldBindJSON(&getProductsQuery); err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		products, err := handler.Handle(&getProductsQuery)
		if err != nil {
			custom_errors.Handle(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, products)
	}
}
