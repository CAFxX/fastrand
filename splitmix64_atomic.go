package fastrand

import (
	"sync/atomic"
)

// AtomicSplitMix64 implements the Java 8 SplittableRandom generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type AtomicSplitMix64 struct {
	SplitMix64
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicSplitMix64) Seed(s uint64) {
	var t SplitMix64
	t.Seed(s)
	atomic.StoreUint64(&r.state, t.state)
}

// Uint64 returns a random uint64.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicSplitMix64) Uint64() uint64 {
	z := atomic.AddUint64(&r.state, 0x9e3779b97f4a7c15)
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
