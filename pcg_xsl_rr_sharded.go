package fastrand

import (
	"runtime"
	"unsafe"
)

// ShardedPCGXSLRR implements the PCG-XSL-RR generator with per-thread (per-P) states.
//
// This generator is safe for concurrent use by multiple goroutines.
// The zero value is a valid state, but it uses a static, all zero seed: use NewShardedPCGXSLRR to instantiate a ShardedPCGXSLRR with a random seed.
type ShardedPCGXSLRR struct {
	states   []paddedPCGXSLRR
	fallback AtomicPCGXSLRR
}

type paddedPCGXSLRR struct {
	PCGXSLRR
	_ [cacheline - unsafe.Sizeof(PCGXSLRR{})%cacheline]byte
}

// NewShardedPCGXSLRR creates a valid ShardedPCGXSLRR instance seeded using crypto/rand.
//
// Increasing the value of GOMAXPROCS after instantiation will likely yield sub-optimal performance.
func NewShardedPCGXSLRR() *ShardedPCGXSLRR {
	p := &ShardedPCGXSLRR{
		states: make([]paddedPCGXSLRR, runtime.GOMAXPROCS(0)),
	}
	for i := range p.states {
		p.states[i].Seed(Seed(), Seed())
	}
	p.fallback.Seed(Seed(), Seed())
	return p
}

// Uint64 returns a random uint64.
//
// This function is safe for concurrent use by multiple goroutines.
func (p *ShardedPCGXSLRR) Uint64() uint64 {
	l := len(p.states) // if p is nil, panic before procPin
	id := procPin()

	if fastrand_nounsafe || l <= id {
		procUnpin()
		return p.fallback.Uint64()
	}

	n := p.states[id].Uint64()
	procUnpin()
	return n
}
