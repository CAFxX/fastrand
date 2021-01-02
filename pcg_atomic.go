package fastrand

import (
	"math/bits"
	"sync/atomic"
)

type AtomicPCG struct {
	PCG
}

func (r *AtomicPCG) Seed(s uint64) {
	atomic.StoreUint64(&r.state, (s+pcgIncr)*pcgMult+pcgIncr)
}

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
