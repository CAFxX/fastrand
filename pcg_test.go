package fastrand

import "testing"

var expPCG = []uint32{0xd386fb93, 0xd1bfe43e, 0xc391d32f, 0xc87f4598, 0x8f712c5c, 0xbb219c74, 0xb8e4973d, 0x8440776c}

func TestPCG(t *testing.T) {
	var r PCG
	r.Seed(1)
	for i, e := range expPCG {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestAtomicPCG(t *testing.T) {
	var r AtomicPCG
	r.Seed(1)
	for i, e := range expPCG {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCGFallback(t *testing.T) {
	var r ShardedPCG // no shards created
	r.fallback.Seed(1)
	for i, e := range expPCG {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCG(t *testing.T) {
	r := NewShardedPCG()
	id := procPin()
	defer procUnpin()
	if fastrand_nounsafe {
		r.fallback.Seed(1)
	} else {
		r.states[id].Seed(1)
	}
	for i, e := range expPCG {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}
