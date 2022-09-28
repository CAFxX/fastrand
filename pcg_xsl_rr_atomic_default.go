//go:build !amd64v2

package fastrand

import (
	"runtime"
	"sync/atomic"

	"github.com/CAFxX/atomics"
	"github.com/klauspost/cpuid/v2"
)

func (r *AtomicPCGXSLRR) cas128(oh, ol, nh, nl uint64) (h, l uint64) {
	p := r.pstate()
	if runtime.GOARCH == "amd64" && cpuid.CPU.Has(cpuid.CX16) {
		return atomics.CompareAndSwap2xUint64(&p[0], oh, ol, nh, nl)
	}
	// fallback implementation
	s := r.plock()
	if !atomic.CompareAndSwapUint64(s, 0, 1) {
		for {
			if atomic.LoadUint64(s) == 0 && atomic.CompareAndSwapUint64(s, 0, 1) {
				break
			}
			cpuYield(1)
		}
	}
	ch, cl := p[0], p[1]
	if ch != oh || cl != ol {
		atomic.StoreUint64(s, 0)
		return ch, cl
	}
	p[0], p[1] = nh, nl
	atomic.StoreUint64(s, 0)
	return nh, nl
}
