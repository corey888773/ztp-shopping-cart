package checkout

import (
	"fmt"
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
)

func Checkout(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cmd Command
		if err := ctx.ShouldBindJSON(&cmd); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := handler.Handle(&cmd)
		if err != nil {
			fmt.Printf("Error handling command: %v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
