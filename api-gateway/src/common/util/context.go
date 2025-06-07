package util

import (
	"context"
	"fmt"
	"log"
)

func CancelWithPanic(cancel context.CancelFunc, err error) {
	cancel()
	fmt.Println("Shutting down the API Gateway...", err)
	log.Fatalf("Error: %v", err)
}
