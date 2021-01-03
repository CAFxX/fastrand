package fastrand

import (
	"math/bits"
	"sync/atomic"
)

// AtomicPCG implements the PCG-XSH-RR generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type AtomicPCG struct {
	PCG
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCG) Seed(s uint64) {
	atomic.StoreUint64(&r.state, (s+pcgIncr)*pcgMult+pcgIncr)
}

// Uint32 returns a random uint32.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCG) Uint32() uint32 {
	i := uint32(0)
	for {
		x := atomic.LoadUint64(&r.state)
		new := x*pcgMult + pcgIncr
		if atomic.CompareAndSwapUint64(&r.state, x, new) {
			x = x ^ (x >> 18)
			count := x >> 59
			return bits.RotateLeft32((uint32)(x>>27), (int)(32-count))
		}
		i += 30
		cpuYield(i)
	}
}
