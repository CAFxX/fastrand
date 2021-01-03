package fastrand

import (
	"runtime"
	"unsafe"
)

const cacheline = 64

// ShardedSplitMix64 implements the Java 8 SplittableRandom generator with per-thread (per-P) states.
//
// It is safe for concurrent use by multiple goroutines.
// The zero value is a valid state, but it uses a static all zero seed: use NewShardedSplitMix64 to instantiate a ShardedSplitMix64 with a random seed.
type ShardedSplitMix64 struct {
	states   []paddedSplitMix64
	fallback AtomicSplitMix64
}

type paddedSplitMix64 struct {
	SplitMix64
	_ [cacheline - unsafe.Sizeof(SplitMix64{})%cacheline]byte
}

// NewShardedSplitMix64 creates a valid ShardedSplitMix64 instance seeded using crypto/rand.
//
// Increasing the value of GOMAXPROCS after instantiation will likely yield sub-optimal performance.
func NewShardedSplitMix64() *ShardedSplitMix64 {
	r := &ShardedSplitMix64{
		states: make([]paddedSplitMix64, runtime.GOMAXPROCS(0)),
	}
	for i := range r.states {
		r.states[i].Seed(seed())
	}
	r.fallback.Seed(seed())
	return r
}

// Uint64 returns a random uint64.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *ShardedSplitMix64) Uint64() uint64 {
	l := len(r.states) // if r is nil, panic before procPin
	id := procPin()

	if l <= id {
		procUnpin()
		return r.fallback.Uint64()
	}

	n := r.states[id].Uint64()
	procUnpin()
	return n
}
