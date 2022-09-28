package fastrand

import (
	"unsafe"
)

// AtomicPCGXSLRR implements the PCG-XSL-RR generator with atomic state updates.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type AtomicPCGXSLRR struct {
	state [4]uint64
}

func (r *AtomicPCGXSLRR) pstate() *[2]uint64 {
	ptr := (uintptr(unsafe.Pointer(&r.state[0]))+15)&^15 + 0
	return (*[2]uint64)(unsafe.Pointer(ptr))
}

func (r *AtomicPCGXSLRR) plock() *uint64 {
	ptr := (uintptr(unsafe.Pointer(&r.state[0]))+15)&^15 + 16
	return (*uint64)(unsafe.Pointer(ptr))
}

// Seed initializes the state with the provided seed.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSLRR) Seed(h, l uint64) {
	var t PCGXSLRR
	t.Seed(h, l)
	nh, nl := t.state[0], t.state[1]
	oh, ol := r.pstate()[0], r.pstate()[1]
	for {
		h, l := r.cas128(oh, ol, nh, nl)
		if h == nh && l == nl {
			return
		}
		oh, ol = h, l
		cpuYield(1)
	}
}

// Uint64 returns a random uint64.
//
// This function is safe for concurrent use by multiple goroutines.
func (r *AtomicPCGXSLRR) Uint64() uint64 {
	oh, ol := r.pstate()[0], r.pstate()[1]
	for {
		nh, nl, n := pcgxslrrStep(oh, ol)
		h, l := r.cas128(oh, ol, nh, nl)
		if h == nh && l == nl {
			return n
		}
		oh, ol = h, l
		cpuYield(1)
	}
}
