package main

import (
	"log"

	"github.com/jay-1409/atimer/internals/timer"
)

func main() {
	log.Println("Starting atimer...")
	handler := timer.NewTimerEventHandler("main-handler", 1024, 4)
	handler.Handler() // Start worker goroutines

	t := timer.NewTimer(4, 1024, handler)
	log.Printf("Initialized timer with %d heaps and handler %s", len(t.Heaps), handler.ID)
}
