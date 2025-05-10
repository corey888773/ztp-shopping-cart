package app

import (
	"context"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common"
)

func CreateApp(appCtx context.Context, config common.Config) (*common.Srv, error) {
	// Setup API server
	server, err := common.NewServer(config)
	if err != nil {
		return nil, err
	}
	server.SetupRouter()
	return server, nil
}
