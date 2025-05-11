package unlock_product

import (
	"fmt"
	"net/http"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/commands"
	"github.com/gin-gonic/gin"
)

func UnlockProduct(handler commands.Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		command := Command{
			ProductID: id,
		}
		fmt.Printf("Command: %+v\n", command)
		err := handler.Handle(&command)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	}
}
