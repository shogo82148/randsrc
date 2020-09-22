/*
Package xoroshiro128p implements xoroshiro128+ 1.0, our best and fastest small-state generator
for floating-point numbers. We suggest to use its upper bits for
floating-point generation, as it is slightly faster than
xoroshiro128++/xoroshiro128**. It passes all tests we are aware of
except for the four lower bits, which might fail linearity tests (and
just those), so if low linear complexity is not considered an issue (as
it is usually the case) it can be used to generate 64-bit outputs, too;
moreover, this generator has a very mild Hamming-weight dependency
making our test (http://prng.di.unimi.it/hwd.php) fail after 5 TB of
output; we believe this slight bias cannot affect any application. If
you are concerned, use xoroshiro128++, xoroshiro128** or xoshiro256+.

We suggest to use a sign test to extract a random Boolean value, and
right shifts to extract subsets of bits.

The state must be seeded so that it is not everywhere zero. If you have
a 64-bit seed, we suggest to seed a splitmix64 generator and use its
output to fill s.

NOTE: the parameters (a=24, b=16, b=37) of this version give slightly
better results in our test than the 2016 version (a=55, b=14, c=36).
*/
package xoroshiro128p

import (
	"math/bits"
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xoroshiro128+ algorithm.
//
// Go port of http://prng.di.unimi.it/xoroshiro128plus.c
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
	result := s0 + s1
	s1 ^= s0
	s0 = bits.RotateLeft64(s0, 24) ^ s1 ^ (s1 << 16) // a, b
	s1 = bits.RotateLeft64(s1, 37)                   // c
	s.state = [2]uint64{s0, s1}
	return result
}
