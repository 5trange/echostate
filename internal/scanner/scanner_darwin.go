// internal/scanner/scanner_darwin.go
//go:build darwin

package scanner

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreWLAN
#import <Foundation/Foundation.h>
#import <CoreWLAN/CoreWLAN.h>

// Structure to hold all WiFi metrics at once
typedef struct {
    long rssi;
    long noise;
    double txRate;
    long channel;
} WiFiMetrics;

WiFiMetrics get_wifi_metrics(){
	// Initialize an empty struct with zeros
    WiFiMetrics metrics = {0, 0, 0.0, 0};

	@autoreleasepool {
		// Get a hold of the shared WiFi client.
		CWWiFiClient *wifiClient = [CWWiFiClient sharedWiFiClient];
		// Get the current Wi-Fi interface.
		CWInterface *interface = [wifiClient interface];

		if(interface != nil) {
			// Get all the data from the interface and populate the metrics struct.
			metrics.rssi = [interface rssiValue];
			metrics.noise = [interface noiseMeasurement];
			metrics.txRate = [interface transmitRate];

			// Try to get the channel number from the interface.
			CWChannel *channel = [interface wlanChannel];
			if(channel != nil) {
				metrics.channel = [channel channelNumber];
			}
		}
	}
	return metrics;
}
*/
import "C"

import (
	"fmt"
	"time"

	"github.com/5trange/echostate/internal/models"
)

type DarwinScanner struct {
}

func (s *DarwinScanner) GetReading() (*models.RFReading, error) {
	// Darwin-specific implementation of the GetReading function
	metrics := C.get_wifi_metrics()
	rssi := int(metrics.rssi) // Get the RSSI value from the metrics struct.

	// Check if RSSI is 0. If it is 0, WiFi interface is likely disconnected.
	if rssi == 0 {
		return nil, fmt.Errorf("WiFi not connected!")
	}

	return &models.RFReading{
		Timestamp: time.Now().UnixMilli(),
		RSSI:      rssi,
		Noise:     int(metrics.noise),
		SNR:       rssi - int(metrics.noise), // Get the SNR value by minusing noise from RSSI.
		TxRate:    float64(metrics.txRate),
		Channel:   int(metrics.channel),
	}, nil
}

func NewScanner() WiFiScanner {
	return &DarwinScanner{}
}
