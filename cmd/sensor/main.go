package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/5trange/echostate/internal/scanner"
)

//go:embed banner.txt
var banner string

func main() {
	// The main method will ask for a scanner and start reading from it.
	// But before we begin, some ART!
	fmt.Println(banner)

	// Create a new WiFi scanner
	wifiScanner := scanner.NewScanner()

	// Poll the scanner at a regular interval (twice every second in this case)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	// Set up a channel to listen for signals
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	// Read RSSI values from the scanner at regular intervals
	for {
		select {
		case <-sigChannel:
			return // Exit the loop if a signal is received

		case <-ticker.C:
			reading, err := wifiScanner.GetReading() // Read the value from the scanner

			// If there was an error, log it and skip to the next reading
			if err != nil {
				log.Printf("Missed reading: %v", err)
				continue
			}

			fmt.Printf("\rCurrent Signal Strength: %d dBm   ", reading.RSSI)
		}
	}
}
