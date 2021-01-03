package fastrand

import (
	"math/rand"
	"sync"
	"testing"

	valyala_fastrand "github.com/valyala/fastrand"
)

func BenchmarkSplitMix64(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r SplitMix64
		r.Seed(Seed())
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkAtomicSplitMix64(b *testing.B) {
	var r AtomicSplitMix64
	r.Seed(Seed())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkShardedSplitMix64(b *testing.B) {
	r := NewShardedSplitMix64()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkPCG(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r PCG
		r.Seed(Seed())
		for pb.Next() {
			use32(r.Uint32())
		}
	})
}

func BenchmarkAtomicPCG(b *testing.B) {
	var r AtomicPCG
	r.Seed(Seed())
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use32(r.Uint32())
		}
	})
}

func BenchmarkShardedPCG(b *testing.B) {
	r := NewShardedPCG()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use32(r.Uint32())
		}
	})
}

func BenchmarkXoshiro256StarStar(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := &Xoshiro256StarStar{}
		r.safeSeed()
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkShardedXoshiro256StarStar(b *testing.B) {
	r := NewShardedXoshiro256StarStar()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkMathRand(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(0))
		for pb.Next() {
			use64(r.Uint64())
		}
	})
}

func BenchmarkMathRandMutex(b *testing.B) {
	var m sync.Mutex
	r := rand.New(rand.NewSource(0))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m.Lock()
			use64(r.Uint64())
			m.Unlock()
		}
	})
}

func BenchmarkValyalaFastrand(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			use32(valyala_fastrand.Uint32())
		}
	})
}

//go:noinline
func use32(uint32) {}

//go:noinline
func use64(uint64) {}
