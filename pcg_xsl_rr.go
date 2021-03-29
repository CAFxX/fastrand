package fastrand

import (
	"math/bits"
)

// PCGXSLRR implements the PCG-XSL-RR generator.
// This generator is not safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type PCGXSLRR struct {
	state [2]uint64 // [0] -> high, [1] -> low
}

const (
	pcgXslMultH = 2549297995355413924
	pcgXslMultL = 4865540595714422341
	pcgXslIncrH = 6364136223846793005
	pcgXslIncrL = 1442695040888963407
)

// Seed initializes the state with the provided seed.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *PCGXSLRR) Seed(h, l uint64) {
	r.state[0], r.state[1] = add128(h, l, pcgXslIncrH, pcgXslIncrL)
	r.Uint64() // as done by the original C implementation
}

// Uint64 returns a random uint64.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *PCGXSLRR) Uint64() (n uint64) {
	r.state[0], r.state[1], n = pcgxslrrStep(r.state[0], r.state[1])
	return
}

func pcgxslrrStep(oh, ol uint64) (nh, nl, n uint64) {
	// output = rotate64(uint64_t(state ^ (state >> 64)), state >> 122);
	n = bits.RotateLeft64(ol^oh, (int)(oh>>(122-64)))
	ph, pl := mul128(oh, ol, pcgXslMultH, pcgXslMultL)
	nh, nl = add128(ph, pl, pcgXslIncrH, pcgXslIncrL)
	return
}

func mul128(ah, al, bh, bl uint64) (h, l uint64) {
	h, l = bits.Mul64(al, bl)
	h += ah * bl
	h += al * bh
	return
}

func add128(ah, al, bh, bl uint64) (h, l uint64) {
	l, h = bits.Add64(al, bl, 0)
	h += ah
	h += bh
	return
}
