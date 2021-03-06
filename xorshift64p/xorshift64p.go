package xorshift64p

import (
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorshift64+ algorithm.
type Source struct {
	a, b uint64
}

// New create a new source.
func New(a, b uint64) *Source {
	if a == 0 && b == 0 {
		a = 1
	}
	return &Source{
		a: a,
		b: b,
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
	s.a = src.Uint64()
	s.b = src.Uint64()
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	a, b := s.a, s.b
	s.a = b
	a ^= a << 23
	a ^= a >> 17
	a ^= b ^ (b >> 26)
	s.b = a
	return a + b
}
