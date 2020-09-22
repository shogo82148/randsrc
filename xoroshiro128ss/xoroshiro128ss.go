/*
Package xoroshiro128ss implements xoroshiro128** 1.0, one of our all-purpose, rock-solid,
small-state generators. It is extremely (sub-ns) fast and it passes all
tests we are aware of, but its state space is large enough only for
mild parallelism.

For generating just floating-point numbers, xoroshiro128+ is even
faster (but it has a very mild bias, see notes in the comments).

The state must be seeded so that it is not everywhere zero. If you have
a 64-bit seed, we suggest to seed a splitmix64 generator and use its
output to fill s.
*/
package xoroshiro128ss

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoroshiro128** algorithm.
//
// Go port of http://prng.di.unimi.it/xoroshiro128starstar.c
type Source struct {
	state [2]uint64
}

// New create a new source.
func New(state [2]uint64) *Source {
	var zero [2]uint64
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
	s0, s1 := s.state[0], s.state[1]
	result := bits.RotateLeft64(s0*5, 7) * 9
	s1 ^= s0
	s0 = bits.RotateLeft64(s0, 24) ^ s1 ^ (s1 << 16) // a, b
	s1 = bits.RotateLeft64(s1, 37)                   // c
	s.state = [2]uint64{s0, s1}
	return result
}
