// internal/scanner/scanner.go
package scanner

import "github.com/5trange/echostate/internal/models"

type WiFiScanner interface {
	GetReading() (*models.RFReading, error)
}
