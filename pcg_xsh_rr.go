package fastrand

import (
	"math/bits"
)

// PCGXSHRR implements the PCG-XSH-RR generator.
// This generator is not safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type PCGXSHRR struct {
	state uint64
}

const (
	pcgMult = 6364136223846793005
	pcgIncr = 1442695040888963407
)

// Seed initializes the state with the provided seed.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *PCGXSHRR) Seed(s uint64) {
	r.state = s + pcgIncr
	r.Uint32() // as done by the original C implementation
}

// Uint32 returns a random uint32.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *PCGXSHRR) Uint32() uint32 {
	x := r.state
	x = x ^ (x >> 18)
	n := bits.RotateLeft32((uint32)(x>>27), (int)(32-(x>>59)))
	r.state = x*pcgMult + pcgIncr
	return n
}
