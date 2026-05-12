package main

import (
	_ "embed"
	"fmt"

	"github.com/5trange/echostate/internal/scanner"
)

//go:embed banner.txt
var banner string

func main() {
	// The main method will ask for a scanner and start reading from it.
	// But before we begin, some ART!
	fmt.Println(banner)

	wifiScanner := scanner.NewScanner()
	rssi, err := wifiScanner.GetReading()

	if err == nil {
		fmt.Println("Timestamp: ", rssi.Timestamp)
		fmt.Println("RSSI Reading: ", rssi.RSSI)
	} else {
		fmt.Println("Error: ", err)
	}
}
