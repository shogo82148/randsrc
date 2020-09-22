package crypto

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using crypto/rand.
type Source struct {
	index int
	buf   [1024]byte
}

// New create a new source.
func New() *Source {
	return &Source{}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() >> 1
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	// do nothing
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	if s.index == 0 {
		_, err := crand.Read(s.buf[:])
		if err != nil {
			panic(err)
		}
	}
	result := binary.BigEndian.Uint64(s.buf[s.index:])
	s.index = (s.index + 8) % len(s.buf)
	return result
}
