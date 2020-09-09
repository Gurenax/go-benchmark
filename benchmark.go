package benchmark

import (
	"fmt"
	"time"
)

// Benchmark struct
type Benchmark struct {
	startTime time.Time
	endTime   time.Time
	duration  time.Duration
}

// Start a new benchmark
func (b *Benchmark) startBenchmark() {
	b.startTime = time.Now()
}

// End the benchmark then record the duration
func (b *Benchmark) endBenchmark() {
	b.endTime = time.Now()
	b.duration = b.endTime.Sub(b.startTime)
}

// Display the recorded duration in milliseconds
func (b *Benchmark) displayResultInMilliseconds(label string) {
	durationInMs := float64(b.duration.Nanoseconds()) * 0.000001
	fmt.Printf("%v %fms\n", label, durationInMs)
}

// Display the recorded duration in seconds
func (b *Benchmark) displayResultInSeconds(label string) {
	durationInSecs := float64(b.duration.Nanoseconds()) * 0.000000001
	fmt.Printf("%v %fseconds\n", label, durationInSecs)
}
