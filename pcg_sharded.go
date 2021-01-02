package fastrand

import (
	"runtime"
	"unsafe"
)

type ShardedPCG struct {
	states   []paddedPCG
	fallback PCG
}

type paddedPCG struct {
	PCG
	_ [cacheline - unsafe.Sizeof(PCG{})%cacheline]byte
}

func NewShardedPCG() *ShardedPCG {
	p := &ShardedPCG{
		states: make([]paddedPCG, runtime.GOMAXPROCS(0)),
	}
	for i := range p.states {
		p.states[i].Seed(seed())
	}
	p.fallback.Seed(seed())
	return p
}

func (p *ShardedPCG) Uint32() uint32 {
	l := len(p.states)
	id := procPin()

	if l <= id {
		procUnpin()
		return p.fallback.Uint32()
	}

	n := p.states[id].Uint32()
	procUnpin()
	return n
}
