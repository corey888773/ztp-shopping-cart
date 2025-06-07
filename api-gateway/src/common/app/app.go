package app

import (
	"context"

	"github.com/corey888773/ztp-shopping-cart/api-gateway/src/common"
)

func CreateApp(ctx context.Context, cfg common.Config) (*common.Srv, error) {
	server, err := common.NewServer(cfg)
	if err != nil {
		return nil, err
	}
	server.SetupRouter()
	return server, nil
}
