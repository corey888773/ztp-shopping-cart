package lock_product

import (
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
)

func LockProduct(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		command := Command{
			ProductID: id,
		}
		err := handler.Handle(&command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
