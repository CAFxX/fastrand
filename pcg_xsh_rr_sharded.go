package fastrand

import (
	"runtime"
	"unsafe"
)

// ShardedPCGXSHRR implements the PCG-XSH-RR generator with per-thread (per-P) states.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state, but it uses a static, all zero seed: use NewShardedPCG to instantiate a ShardedPCG with a random seed.
type ShardedPCGXSHRR struct {
	states   []paddedPCGXSHRR
	fallback AtomicPCGXSHRR
}

type paddedPCGXSHRR struct {
	PCGXSHRR
	_ [cacheline - unsafe.Sizeof(PCGXSHRR{})%cacheline]byte
}

// NewShardedPCGXSHRR creates a valid ShardedPCG instance seeded using crypto/rand.
//
// Increasing the value of GOMAXPROCS after instantiation will likely yield sub-optimal performance.
func NewShardedPCGXSHRR() *ShardedPCGXSHRR {
	p := &ShardedPCGXSHRR{
		states: make([]paddedPCGXSHRR, runtime.GOMAXPROCS(0)),
	}
	for i := range p.states {
		p.states[i].Seed(Seed())
	}
	p.fallback.Seed(Seed())
	return p
}

// Uint32 returns a random uint32.
//
// This function is safe for concurrent use by multiple goroutines.
func (p *ShardedPCGXSHRR) Uint32() uint32 {
	l := len(p.states) // if p is nil, panic before procPin
	id := procPin()

	if fastrand_nounsafe || l <= id {
		procUnpin()
		return p.fallback.Uint32()
	}

	n := p.states[id].Uint32()
	procUnpin()
	return n
}
