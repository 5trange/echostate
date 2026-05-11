// internal/scanner/scanner_darwin.go
//go:build darwin

package scanner

/*
#cgo CFLAGS: -x objective-c
int fun(){
	return 1;
}
*/
import "C"

import (
	"github.com/5trange/echostate/internal/models"
)

type DarwinScanner struct {
}

func (s *DarwinScanner) GetReading() (*models.RFReading, error) {
	// Darwin-specific implementation of the GetReading function.
	return nil, nil // Placeholder
}

func NewScanner() WiFiScanner {
	return &DarwinScanner{}
}
