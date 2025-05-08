package util

import (
	"context"
	"fmt"
	"log"
	"time"
)

func CreateContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func CancelWithPanic(cancel context.CancelFunc, err error) {
	cancel()
	fmt.Println("Shutting down the application...", err)
	log.Fatalf("Error: %v", err)
}
