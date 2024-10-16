package main

import (
	"HomeWork_2/internal"
	"log"
)

func main() {
	internal.RunServer()
	if err := internal.RunServer; err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
