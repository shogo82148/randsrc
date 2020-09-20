package xoshiro256ss

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/xorshift64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoshiro256** algorithm.
//
// http://prng.di.unimi.it/
// http://prng.di.unimi.it/xoshiro256starstar.c
type Source struct {
	state [4]uint64
}

// New create a new source.
func New(state [4]uint64) *Source {
	var zero [4]uint64
	if state == zero {
		state[0] = 1
	}
	return &Source{
		state: state,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() & 0x7FFFFFFFFFFFFFFF
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	src := xorshift64.New(uint64(seed))
	for i := range s.state {
		s.state[i] = src.Uint64()
	}
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	result := bits.RotateLeft64(s.state[1]*5, 7) * 9
	t := s.state[1] << 17
	s.state[2] ^= s.state[0]
	s.state[3] ^= s.state[1]
	s.state[1] ^= s.state[2]
	s.state[0] ^= s.state[3]
	s.state[2] ^= t
	s.state[3] = bits.RotateLeft64(s.state[3], 45)
	return result
}
