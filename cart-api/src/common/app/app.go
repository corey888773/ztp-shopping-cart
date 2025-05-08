package app

import (
	"context"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/util"
)

func CreateApp(appCtx context.Context, config util.Config) (*common.Srv, error) {
	// Setup API server
	server, err := common.NewServer(config)
	if err != nil {
		return nil, err
	}
	server.SetupRouter()
	return server, nil
}
