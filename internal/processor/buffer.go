// internal/processor/buffer.go
package processor

type RSSIRingBuffer struct {
	rssiReading []int
	size        int
	head        int
	count       int
}

func NewRSSIRingBuffer(size int) *RSSIRingBuffer {
	return &RSSIRingBuffer{
		rssiReading: make([]int, size),
		size:        size,
		head:        0, // Start from index 0
		count:       0,
	}
}

func (b *RSSIRingBuffer) Insert(rssi int) {
	b.rssiReading[b.head] = rssi

	// (head + 1) % size wraps around to 0 when head reaches the end
	b.head = (b.head + 1) % b.size

	if b.count < b.size {
		b.count++
	}
}
