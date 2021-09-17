package fastrand

import (
	"github.com/CAFxX/atomic128"
)

// AtomicPCGXSLRR implements the PCG-XSL-RR generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
// This structure must not be copied.
type AtomicPCGXSLRR struct {
	state atomic128.Uint128
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSLRR) Seed(h, l uint64) {
	var t PCGXSLRR
	t.Seed(h, l)
	atomic128.StoreUint128(&r.state, t.state)
}

// Uint64 returns a random uint64.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSLRR) Uint64() uint64 {
	for {
		s := atomic128.LoadUint128(&r.state)
		nh, nl, n := pcgxslrrStep(s[0], s[1])
		if atomic128.CompareAndSwapUint128(&r.state, s, [2]uint64{nh, nl}) {
			return n
		}
		cpuYield(1)
	}
}
