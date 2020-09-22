package xorwow

import (
	"math/rand"

	"github.com/shogo82148/randsrc/splitmix64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorwow algorithm.
// https://en.wikipedia.org/wiki/Xorshift#xorwow
// Marsaglia, George (July 2003). "Xorshift RNGs". Journal of Statistical Software. 8 (14).
type Source struct {
	a, b, c, d, e uint32
	counter       uint32
}

// New create a new source.
func New(a, b, c, d, e uint32) *Source {
	if a == 0 && b == 0 && c == 0 && d == 0 && e == 0 {
		a = 1
	}
	return &Source{
		a: a,
		b: b,
		c: c,
		d: d,
		e: e,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := int64(s.Uint32()) << 31
	x += int64(s.Uint32()) >> 1
	return x
}

// Uint32 returns pseudo-random uint32 values in the range [0, 1<<32).
func (s *Source) Uint32() uint32 {
	t, u := s.e, s.a
	s.e, s.d, s.c, s.b = s.d, s.c, s.b, s.a
	t ^= t >> 2
	t ^= t << 1
	t ^= u ^ (u << 4)
	s.a = t
	s.counter += 362437
	return t + s.counter
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	src := splitmix64.New(uint64(seed))
	x := src.Uint64()
	y := src.Uint64()
	z := src.Uint64()
	s.a = uint32(x >> 32)
	s.b = uint32(x)
	s.c = uint32(y >> 32)
	s.d = uint32(y)
	s.e = uint32(z >> 32)
	s.counter = uint32(z)
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := uint64(s.Uint32()) << 32
	x += uint64(s.Uint32())
	return x
}
