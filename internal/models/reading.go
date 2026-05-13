// internal/models/reading.go
package models

// Used RFReading to export the struct to be used in other packages.
// RFReading represents a single reading from a sensor?
type RFReading struct {
	Timestamp int64   `json:"timestamp"`
	RSSI      int     `json:"rssi"`
	Noise     int     `json:"noise"`   // Background radiation
	SNR       int     `json:"snr"`     // Signal-to-noise ratio
	TxRate    float64 `json:"tx_rate"` // Transaction rate in Mbps
	Channel   int     `json:"channel"` // The channel we are on
}
