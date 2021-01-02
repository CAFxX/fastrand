package fastrand

import (
	"crypto/rand"
	"encoding/binary"
)

func seed() uint64 {
	b := [8]byte{}
	n, err := rand.Read(b[:])
	if n != 8 || err != nil {
		panic("unable to read seed")
	}
	return binary.LittleEndian.Uint64(b[:])
}
