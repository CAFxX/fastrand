package fastrand

import (
	"runtime"
	"unsafe"
)

const cacheline = 64

type ShardedSplitMix64 struct {
	states   []paddedSplitMix64
	fallback AtomicSplitMix64
}

type paddedSplitMix64 struct {
	SplitMix64
	_ [cacheline - unsafe.Sizeof(SplitMix64{})%cacheline]byte
}

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
