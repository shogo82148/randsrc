/*
Package xoshiro512p implements xoshiro512+ 1.0, our generator for floating-point numbers with
increased state size. We suggest to use its upper bits for
floating-point generation, as it is slightly faster than xoshiro512**.
It passes all tests we are aware of except for the lowest three bits,
which might fail linearity tests (and just those), so if low linear
complexity is not considered an issue (as it is usually the case) it
can be used to generate 64-bit outputs, too.

We suggest to use a sign test to extract a random Boolean value, and
right shifts to extract subsets of bits.

The state must be seeded so that it is not everywhere zero. If you have
a 64-bit seed, we suggest to seed a splitmix64 generator and use its
output to fill s.
*/
package xoshiro512p

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoshiro512+ algorithm.
//
// Go port of http://prng.di.unimi.it/xoshiro512plus.c
type Source struct {
	state [8]uint64
}

// New create a new source.
func New(state [8]uint64) *Source {
	var zero [8]uint64
	if state == zero {
		state[0] = 1
	}
	return &Source{
		state: state,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() >> 1
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
	result := s.state[0] + s.state[2]
	t := s.state[1] << 11
	s.state[2] ^= s.state[0]
	s.state[5] ^= s.state[1]
	s.state[1] ^= s.state[2]
	s.state[7] ^= s.state[3]
	s.state[3] ^= s.state[4]
	s.state[4] ^= s.state[5]
	s.state[0] ^= s.state[6]
	s.state[6] ^= s.state[7]
	s.state[6] ^= t
	s.state[7] = bits.RotateLeft64(s.state[7], 21)
	return result
}
