/*
Package xoroshiro1024s implements xoroshiro1024* 1.0, our large-state generator for
floating-point numbers. We suggest to use its upper bits for
floating-point generation, as it is slightly faster than
xoroshiro1024++/xoroshiro1024**.  Its state however is too large--in
general, the xoshiro256 family should be preferred. It is
a better replacement for xorshift1024*.

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
package xoroshiro1024s

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoroshiro1024* algorithm.
//
// Go port of http://prng.di.unimi.it/xoroshiro1024star.c
type Source struct {
	state [16]uint64
	p     int
}

// New create a new source.
func New(state [16]uint64) *Source {
	var zero [16]uint64
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
	q := s.p
	p := (s.p + 1) % 16
	s0, s15 := s.state[p], s.state[q]
	result := s0 * 0x9e3779b97f4a7c13

	s15 ^= s0
	s.state[q] = bits.RotateLeft64(s0, 25) ^ s15 ^ (s15 << 27)
	s.state[p] = bits.RotateLeft64(s15, 36)
	s.p = p
	return result
}
