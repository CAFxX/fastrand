package fastrand

// SplitMix64 implements the Java 8 SplittableRandom generator.
// This generator is not safe for concurrent use by multiple goroutines.
// The zero value is a valid state: Seed() can be called to set a custom seed.
type SplitMix64 struct {
	state uint64
}

// Seed initializes the state with the provided seed.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *SplitMix64) Seed(s uint64) {
	r.state = s
}

// Uint64 returns a random uint64.
//
// This function is not safe for concurrent use by multiple goroutines.
func (r *SplitMix64) Uint64() uint64 {
	r.state += 0x9e3779b97f4a7c15
	z := r.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
