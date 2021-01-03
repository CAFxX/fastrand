// +build !fastrand_nounsafe

package fastrand

import _ "unsafe"

const fastrand_nounsafe = false

//go:linkname procPin runtime.procPin
func procPin() int

//go:linkname procUnpin runtime.procUnpin
func procUnpin()
