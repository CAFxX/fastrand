package fastrand

type SplitMix64 struct {
	state uint64
}

func (r *SplitMix64) Seed(s uint64) {
	r.state = s
}

func (r *SplitMix64) Uint64() uint64 {
	r.state += 0x9e3779b97f4a7c15
	z := r.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
