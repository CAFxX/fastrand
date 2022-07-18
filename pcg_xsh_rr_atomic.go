package fastrand

import (
	"sync/atomic"
)

// AtomicPCGXSHRR implements the PCG-XSH-RR generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type AtomicPCGXSHRR struct {
	PCGXSHRR
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSHRR) Seed(s uint64) {
	var t PCGXSHRR
	t.Seed(s)
	atomic.StoreUint64(&r.state, t.state)
}

// Uint32 returns a random uint32.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSHRR) Uint32() uint32 {
	i := uint32(0)
	for {
		old := atomic.LoadUint64(&r.state)
		var t PCGXSHRR
		t.state = old
		n := t.Uint32()
		if atomic.CompareAndSwapUint64(&r.state, old, t.state) {
			return n
		}
		i += 30
		cpuYield(i)
	}
}
