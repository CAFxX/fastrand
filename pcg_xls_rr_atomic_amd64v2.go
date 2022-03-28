//go:build amd64.v2

package fastrand

import "github.com/CAFxX/atomics"

func (r *AtomicPCGXSLRR) cas128(oh, ol, nh, nl uint64) (h, l uint64) {
	// TODO: when the compiler can guarantee 128-bit alignments, we can drop the third uint64 of state
	p := r.pstate()
	// AMD64=v2 includes CMPXCHG16, so no need to check whether it's available
	return atomics.CompareAndSwap2xUint64(&p[0], oh, ol, nh, nl)
}
