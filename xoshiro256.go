package fastrand

import "math/bits"

// Xoshiro256StarStar implements the Xoshiro256** PRNG.
//
// This generator is not safe for concurrent use by multiple goroutines.
// The zero value is not a valid state: Seed() must be called
// before generating random numbers.
type Xoshiro256StarStar struct {
	state [4]uint64
}

// Uint64 returns a random uint64.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *Xoshiro256StarStar) Uint64() uint64 {
	// TODO: this function has, unfortunately, an inline cost of 81.
	// It would be ideal, especially for the sharded variant, if it was inlineable.

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

// Seed sets the seed for the generator.
//
// The seed should not be all zeros (i.e. at least one of the four
// uint64 should be non-zero).
// This function is not safe for concurrent use by multiple goroutines.
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
