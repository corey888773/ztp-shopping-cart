package get_all_products

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/queries"
	"github.com/corey888773/ztp-shopping-cart/products-api/src/features"
	"github.com/gin-gonic/gin"
)

// GetAllProducts returns a handler func for fetching all products.
func GetAllProducts(handler queries.Handler) gin.HandlerFunc {
	errHandler := features.NewErrorHandler(DomainErrors)

	return func(ctx *gin.Context) {
		products, err := handler.Handle(&Query{})
		if err != nil {
			errHandler.Handle(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, products)
	}
}
