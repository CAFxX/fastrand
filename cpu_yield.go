// +build !fastrand_nounsafe

package fastrand

import _ "unsafe"

//go:linkname cpuYield runtime.procyield
func cpuYield(n uint32)
