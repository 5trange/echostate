// internal/models/reading.go
package models

// Exporting because it's used outside of this package.
// Used RFReading to export the struct to be used in other packages.
// RFReading represents a single reading from a sensor?
type RFReading struct {
	Timestamp int64 `json:"timestamp"`
	RSSI      int   `json:"rssi"`
}
