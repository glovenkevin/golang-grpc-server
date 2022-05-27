package main

import (
	"log"

	"grpc-client/setup"
)

func main() {

	err := setup.Load()
	if err != nil {
		log.Fatalf("Failed to start: %v", err)
	}

	if err := setup.Start(); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
