package main

import (
	"internal"
	"log"
)

func main() {
	internal.RunServer()
	if err := internal.RunServer; err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
