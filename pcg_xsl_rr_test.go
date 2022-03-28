package fastrand

import "testing"

var expPCGXSLRR = []uint64{
	0x7a2e19230dd00811, 0x48e426b893aada61, 0xd7d5af49cb7f615b, 0xa556d783aa38c3c8,
	0xccc30b1f8ff7f3c5, 0xc155276314310b90, 0xa343a0176679feff, 0xf80d9327c9d71a7a,
}

func TestPCGXSLRR(t *testing.T) {
	var r PCGXSLRR
	r.Seed(1, 2)
	for i, e := range expPCGXSLRR {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestAtomicPCGXSLRR(t *testing.T) {
	var r AtomicPCGXSLRR
	r.Seed(1, 2)
	for i, e := range expPCGXSLRR {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCGXSLRRFallback(t *testing.T) {
	var r ShardedPCGXSLRR // no shards created
	r.fallback.Seed(1, 2)
	for i, e := range expPCGXSLRR {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedPCGXSLRR(t *testing.T) {
	r := NewShardedPCGXSLRR()
	id := procPin()
	defer procUnpin()
	if fastrand_nounsafe {
		r.fallback.Seed(1, 2)
	} else {
		r.states[id].Seed(1, 2)
	}
	for i, e := range expPCGXSLRR {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}
