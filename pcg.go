package fastrand

import (
	"math/bits"
)

type PCG struct {
	state uint64
}

const (
	pcgMult = 6364136223846793005
	pcgIncr = 1442695040888963407
)

func (r *PCG) Seed(s uint64) {
	r.state = s + pcgIncr
	r.Uint32()
}

func (r *PCG) Uint32() uint32 {
	x := r.state
	x = x ^ (x >> 18)
	n := bits.RotateLeft32((uint32)(x>>27), (int)(32-(x>>59)))
	r.state = x*pcgMult + pcgIncr
	return n
}
