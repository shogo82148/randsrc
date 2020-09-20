package xorshift128

import (
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorshift128 algorithm.
//
// Marsaglia, George (July 2003). "Xorshift RNGs". Journal of Statistical Software. 8 (14).
type Source struct {
	a, b, c, d uint32
}

// New create a new source.
func New(state1, state2 uint64) *Source {
	if state1 == 0 && state2 == 0 {
		state1 = 1
	}
	return &Source{
		a: uint32(state1 >> 32),
		b: uint32(state1),
		c: uint32(state2 >> 32),
		d: uint32(state2),
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() & 0x7FFFFFFFFFFFFFFF
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	src := splitmix64.New(uint64(seed))
	x := src.Uint64()
	y := src.Uint64()
	s.a = uint32(x >> 32)
	s.b = uint32(x)
	s.c = uint32(y >> 32)
	s.d = uint32(y)
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	t := s.d
	u := s.a
	s.d, s.c, s.b = s.c, s.b, s.a
	t ^= t << 11
	t ^= t >> 8
	s.a = t ^ u ^ (u >> 19)
	return (uint64(s.b) << 32) + uint64(s.a)
}
