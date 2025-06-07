package features

import (
	"net/http"

	"slices"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/custom_errors"
	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	domainErrors []string
}

func NewErrorHandler(domainErrors []string) *ErrorHandler {
	return &ErrorHandler{
		domainErrors: domainErrors,
	}
}

func (h *ErrorHandler) Handle(ctx *gin.Context, err error) {
	if slices.Contains(h.domainErrors, err.Error()) {
		custom_errors.WithErrorMessage(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	custom_errors.Handle(ctx, err)
}
