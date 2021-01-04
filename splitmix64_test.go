package fastrand

import "testing"

var expSplitMix64 = []uint64{0x910a2dec89025cc1, 0xbeeb8da1658eec67, 0xf893a2eefb32555e, 0x71c18690ee42c90b, 0x71bb54d8d101b5b9, 0xc34d0bff90150280, 0xe099ec6cd7363ca5, 0x85e7bb0f12278575}

func TestSplitMix64(t *testing.T) {
	var r SplitMix64
	r.Seed(1)
	for i, e := range expSplitMix64 {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestAtomicSplitMix64(t *testing.T) {
	var r AtomicSplitMix64
	r.Seed(1)
	for i, e := range expSplitMix64 {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedSplitMix64Fallback(t *testing.T) {
	var r ShardedSplitMix64 // no shards created
	r.fallback.Seed(1)
	for i, e := range expSplitMix64 {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedSplitMix64(t *testing.T) {
	r := NewShardedSplitMix64()
	id := procPin()
	defer procUnpin()
	if fastrand_nounsafe {
		r.fallback.Seed(1)
	} else {
		r.states[id].Seed(1)
	}
	for i, e := range expSplitMix64 {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}
