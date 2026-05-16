// internal/processor/buffer.go
package processor

import "strings"

type SignalBuffer struct {
	reading []int
	size    int
	head    int
	count   int
}

func NewSignalBuffer(size int) *SignalBuffer {
	return &SignalBuffer{
		reading: make([]int, size),
		size:    size,
		head:    0, // Start from index 0
		count:   0,
	}
}

// Inserts a value into the buffer
func (b *SignalBuffer) Insert(rssi int) {
	b.reading[b.head] = rssi

	// (head + 1) % size wraps around to 0 when head reaches the end
	b.head = (b.head + 1) % b.size

	// Increment count if buffer is not full, otherwise it will wrap around
	if b.count < b.size {
		b.count++
	}
}

// This gets the average from the values in the ring buffer.
func (b *SignalBuffer) GetSMA() float64 {
	if b.count == 0 {
		return 0.0
	}
	sum := 0

	for i := 0; i < b.count; i++ {
		sum += b.reading[i]
	}
	// Cast both to float64 because the function returns float64
	return float64(sum) / float64(b.count)
}

func (b *SignalBuffer) GetSparkline() string {
	var sb strings.Builder
	sb.Grow(b.size * 4) // Unicode characters take 4 bytes each, so we multiply the size by 4

	blocks := []rune{' ', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

	startIndex := 0
	if b.count == b.size {
		startIndex = b.head
	}

	// Add the padding so the TUI doesn't look weird when we start
	for range b.size - b.count {
		sb.WriteString(" ")
	}

	for i := 0; i < b.count; i++ {
		readIndex := (startIndex + i) % b.size
		normalizedSNR := normalizeSNR(b.reading[readIndex])
		sb.WriteRune(blocks[normalizedSNR])
	}
	return sb.String()
}
