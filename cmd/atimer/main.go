package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jay-1409/atimer/internals/timer"
)

func main() {
	defaultPort := os.Getenv("PORT")
	if defaultPort == "" {
		defaultPort = "8080"
	}
	portFlag := flag.String("port", defaultPort, "Port to listen on for the HTTP API server")
	flag.Parse()
	log.Println("Starting atimer...")
	t := timer.NewTimer(4, 1024)
	t.Start()
	log.Printf("Initialized timer with %d heaps", len(t.Heaps))
	addr := fmt.Sprintf(":%s", *portFlag)
	if err := t.StartServer(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
	log.Printf("HTTP server succesfully started")
}
