package main

import (
	"fmt"

	"github.com/5trange/echostate/internal/scanner"
)

func main() {
	// The main method will ask for a scanner and start reading from it.
	wifiScanner := scanner.NewScanner()
	rssi, err := wifiScanner.GetReading()

	if err == nil {
		fmt.Println("Timestamp: ", rssi.Timestamp)
		fmt.Println("RSSI Reading: ", rssi.RSSI)
	} else {
		fmt.Println("Error: ", err)
	}
}
