package main

import (
	"log"
	"word-search-in-files/pkg/server"
)

func main() {
	port := "8080"
	log.Printf("Start listening server on port %s...\n", port)
	err := server.StartServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
