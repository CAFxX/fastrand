package fastrand

import "math/bits"

type Xoshiro256StarStar struct {
	state [4]uint64
}

func (r *Xoshiro256StarStar) Uint64() uint64 {
	result := bits.RotateLeft64(r.state[1]*5, 7) * 9

	t := r.state[1] << 17

	r.state[2] ^= r.state[0]
	r.state[3] ^= r.state[1]
	r.state[1] ^= r.state[2]
	r.state[0] ^= r.state[3]

	r.state[2] ^= t

	r.state[3] = bits.RotateLeft64(r.state[3], 45)

	return result
}

func (r *Xoshiro256StarStar) Seed(s0, s1, s2, s3 uint64) {
	r.state[0] = s0
	r.state[1] = s1
	r.state[2] = s2
	r.state[3] = s3
}

func (r *Xoshiro256StarStar) safeSeed() {
retry:
	a, b, c, d := seed(), seed(), seed(), seed()
	if a == 0 && b == 0 && c == 0 && d == 0 {
		goto retry
	}
	r.Seed(a, b, c, d)
}
