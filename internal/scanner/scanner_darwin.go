// internal/scanner/scanner_darwin.go
//go:build darwin

package scanner

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreWLAN
#import "Foundation/Foundation.h"
#import "CoreWLAN/CoreWLAN.h"
long get_rssi(){
	@autoreleasepool {
		// Get a hold of the shared WiFi client.
		CWWiFiClient *wifiClient = [CWWiFiClient sharedWiFiClient];
		// Get the current Wi-Fi interface.
		CWInterface *interface = [wifiClient interface];

		// Check if the interface is nil and return 0 if it is. This indicates no Wi-Fi connection is established.
		if(interface == nil) {
			return 0;
		}

		// Get the current RSSI value and return it.
		return [interface rssiValue];
	}
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
	// Get the current RSSI value using the C function and return it as a RFReading.
	val := C.get_rssi()
	rssi := int(val)

	if rssi == 0 {
		return nil, fmt.Errorf("wi-fi not connected")
	}

	return &models.RFReading{
		Timestamp: time.Now().UnixMilli(),
		RSSI:      rssi,
	}, nil
}

func NewScanner() WiFiScanner {
	return &DarwinScanner{}
}
