package main

import (
	"log"

	"github.com/jay-1409/atimer/internals/timer"
)

func main() {
	log.Println("Starting atimer...")
	t := timer.NewTimer(4, 1024)
	t.Start()
	log.Printf("Initialized timer with %d heaps", len(t.Heaps))

	// Keep the main goroutine alive so the background heap runners and event workers can run
	select {}
}
