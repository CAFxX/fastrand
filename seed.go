package fastrand

import (
	"crypto/rand"
	"encoding/binary"
)

// Seed returns a random uint64 from crypto/rand.
func Seed() uint64 {
	b := [8]byte{}
	n, err := rand.Read(b[:])
	if n != 8 || err != nil {
		panic("unable to read seed")
	}
	return binary.LittleEndian.Uint64(b[:])
}
