package fastrand

import "testing"

var expPCGXSHRR = []uint32{0xd386fb93, 0xd1bfe43e, 0xc391d32f, 0xc87f4598, 0x8f712c5c, 0xbb219c74, 0xb8e4973d, 0x8440776c}

func TestPCGXSHRR(t *testing.T) {
	var r PCGXSHRR
	r.Seed(1)
	for i, e := range expPCGXSHRR {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestAtomicPCGXSHRR(t *testing.T) {
	var r AtomicPCGXSHRR
	r.Seed(1)
	for i, e := range expPCGXSHRR {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCGXSHRRFallback(t *testing.T) {
	var r ShardedPCGXSHRR // no shards created
	r.fallback.Seed(1)
	for i, e := range expPCGXSHRR {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCGXSHRR(t *testing.T) {
	r := NewShardedPCGXSHRR()
	id := procPin()
	defer procUnpin()
	if fastrand_nounsafe {
		r.fallback.Seed(1)
	} else {
		r.states[id].Seed(1)
	}
	for i, e := range expPCGXSHRR {
		if g := r.Uint32(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}
