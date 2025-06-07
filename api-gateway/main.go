package main

import (
	"context"
	"fmt"
	"log"

	"github.com/corey888773/ztp-shopping-cart/api-gateway/src/common"
	"github.com/corey888773/ztp-shopping-cart/api-gateway/src/common/app"
	"github.com/corey888773/ztp-shopping-cart/api-gateway/src/common/util"
)

func main() {
	// Create application context with cancel
	appCtx, appCancel := context.WithCancel(context.Background())
	defer appCancel()
	go func() {
		<-appCtx.Done()
		log.Println("Shutting down API Gateway...")
	}()

	// Load configuration from app.env
	config, err := util.LoadConfig[common.Config](".")
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}

	// Initialize gateway application
	server, err := app.CreateApp(appCtx, config)
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}

	// Start HTTP server
	addr := fmt.Sprintf(":%s", config.ServerPort)
	if err := server.Start(addr); err != nil {
		util.CancelWithPanic(appCancel, err)
	}
}
