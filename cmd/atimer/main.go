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
	heapsFlag := flag.Int("heaps", 4, "Number of concurrent timer heaps")
	workersFlag := flag.Int("workers", 2, "Number of notification workers per heap")
	flag.Parse()

	log.Println("Starting atimer...")
	t := timer.NewTimer(*heapsFlag, 1024, *workersFlag)
	t.Start()
	log.Printf("Initialized timer with %d heaps (workers per heap: %d)", len(t.Heaps), *workersFlag)
	addr := fmt.Sprintf(":%s", *portFlag)
	if err := t.StartServer(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
