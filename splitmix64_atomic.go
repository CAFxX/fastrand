package fastrand

import (
	"sync/atomic"
)

type AtomicSplitMix64 struct {
	SplitMix64
}

func (r *AtomicSplitMix64) Seed(s uint64) {
	atomic.StoreUint64(&r.state, s)
}

func (r *AtomicSplitMix64) Uint64() uint64 {
	z := atomic.AddUint64(&r.state, 0x9e3779b97f4a7c15)
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
