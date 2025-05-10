package main

import (
	"context"
	"fmt"
	"log"

	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/app"
	"github.com/corey888773/ztp-shopping-cart/cart-api/src/common/util"
)

func main() {
	// Create context
	appCtx, appCancel := context.WithCancel(context.Background())
	go func() {
		<-appCtx.Done()
		log.Println("Shutting down the application in 2 seconds...")
	}()

	config, err := util.LoadConfig[common.Config](".")
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}

	// Start API server
	server, err := app.CreateApp(appCtx, config)
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}

	defer server.Stop() // Stop the server when the application is done
	err = server.Start(fmt.Sprintf(":%s", config.ServerPort))
	if err != nil {
		util.CancelWithPanic(appCancel, err)
	}
}
