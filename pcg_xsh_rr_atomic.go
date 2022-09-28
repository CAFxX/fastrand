package fastrand

import (
	"sync/atomic"
)

// AtomicPCGXSHRR implements the PCG-XSH-RR generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type AtomicPCG struct {
	state atomic.Uint64
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSHRR) Seed(s uint64) {
	var t PCGXSHRR
	t.Seed(s)
	r.state.Store(t.state)
}

// Uint32 returns a random uint32.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSHRR) Uint32() uint32 {
	i := uint32(0)
	for {
		old := r.state.Load()
		var t PCGXSHRR
		t.state = old
		n := t.Uint32()
		if r.state.CompareAndSwap(old, t.state) {
			return n
		}
		i += 30
		cpuYield(i)
	}
}
