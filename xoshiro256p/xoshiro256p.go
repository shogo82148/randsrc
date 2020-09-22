/*
Package xoshiro256p implements xoshiro256+ 1.0, our best and fastest generator for floating-point
numbers. We suggest to use its upper bits for floating-point
generation, as it is slightly faster than xoshiro256++/xoshiro256**. It
passes all tests we are aware of except for the lowest three bits,
which might fail linearity tests (and just those), so if low linear
complexity is not considered an issue (as it is usually the case) it
can be used to generate 64-bit outputs, too.

We suggest to use a sign test to extract a random Boolean value, and
right shifts to extract subsets of bits.

The state must be seeded so that it is not everywhere zero. If you have
a 64-bit seed, we suggest to seed a splitmix64 generator and use its
output to fill s.
*/
package xoshiro256p

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoshiro256+ algorithm.
//
// Go port of http://prng.di.unimi.it/xoshiro256plus.c
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
	src := splitmix64.New(uint64(seed))
	for i := range s.state {
		s.state[i] = src.Uint64()
	}
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	result := s.state[0] + s.state[3]
	t := s.state[1] << 17
	s.state[2] ^= s.state[0]
	s.state[3] ^= s.state[1]
	s.state[1] ^= s.state[2]
	s.state[0] ^= s.state[3]
	s.state[2] ^= t
	s.state[3] = bits.RotateLeft64(s.state[3], 45)
	return result
}
