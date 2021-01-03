package fastrand

import (
	"runtime"
	"sync"
	"unsafe"
)

// ShardedXoshiro256StarStar implements the Xoshiro256** PRNG with per-thread (per-P) states.
// It is safe for concurrent use by multiple goroutines.
// The zero value is not a valid state: use NewShardedXoshiro256StarStar to construct a valid instance.
// Increasing the value of GOMAXPROCS after instantiation will likely yield sub-optimal performance.
type ShardedXoshiro256StarStar struct {
	states        []paddedXoshiro256
	fallbackMutex sync.Mutex
	fallback      Xoshiro256StarStar
}

type paddedXoshiro256 struct {
	Xoshiro256StarStar
	_ [cacheline - unsafe.Sizeof(Xoshiro256StarStar{})%cacheline]byte
}

// NewShardedXoshiro256StarStar creates a valid ShardedXoshiro256StarStar instance.
// The instance is seeded using crypto/rand.
// Increasing the value of GOMAXPROCS after instantiation will likely yield sub-optimal performance.
func NewShardedXoshiro256StarStar() *ShardedXoshiro256StarStar {
	r := &ShardedXoshiro256StarStar{}
	r.states = make([]paddedXoshiro256, runtime.GOMAXPROCS(0))
	for i := range r.states {
		r.states[i].Xoshiro256StarStar.safeSeed()
	}
	r.fallback.safeSeed()
	return r
}

// Uint64 returns a random uint64.
// It is safe for concurrent use by multiple goroutines.
func (r *ShardedXoshiro256StarStar) Uint64() uint64 {
	l := len(r.states) // if r is nil, panic before procPin
	id := procPin()

	if l <= id {
		procUnpin()
		r.fallbackMutex.Lock()
		n := r.fallback.Uint64()
		r.fallbackMutex.Unlock()
		return n
	}

	n := r.states[id].Uint64()
	procUnpin()
	return n
}
