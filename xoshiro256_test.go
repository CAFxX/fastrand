package fastrand

import "testing"

var expXoroshiro = []uint64{0x0, 0x1680, 0x1680, 0x2d001680, 0x2d000002d000000, 0x5a5a2d001680, 0x2d05a002d0b4000, 0x1c005a002d000000}

func TestXoshiro256StarStar(t *testing.T) {
	var r Xoshiro256StarStar
	r.Seed(1, 0, 0, 0)
	for i, e := range expXoroshiro {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedXoshiro256StarStarFallback(t *testing.T) {
	var r ShardedXoshiro256StarStar // no shards created
	r.fallback.Seed(1, 0, 0, 0)
	for i, e := range expXoroshiro {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}

func TestShardedXoshiro256StarStar(t *testing.T) {
	r := NewShardedXoshiro256StarStar()
	id := procPin()
	defer procUnpin()
	r.states[id].Seed(1, 0, 0, 0)
	for i, e := range expXoroshiro {
		if g := r.Uint64(); g != e {
			t.Errorf("i=%d, expected=%x, got=%x", i, e, g)
		}
	}
}
