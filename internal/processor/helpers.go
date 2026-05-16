// internal/processor/buffer.go
package processor

// Helper function to draw the sparkline graph
// For a unicode sparkline graph the available characters are
// ▂▃▅▆▇█ 0-7 indexes
func normalizeSNR(snr int) int {
	// Handle the edgecases, do not try to normalize these values
	if snr <= 10 {
		return 0
	} else if snr >= 60 {
		return 7
	}

	// We need to normalize whatever value of SNR this is
	// scaled = ((value - min) * max_index) / (max - min)
	scaled := ((snr - 10) * 7) / 50
	return scaled
}
