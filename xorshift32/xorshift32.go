package xorshift32

import "math/rand"

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorshift32 algorithm.
//
// Marsaglia, George (July 2003). "Xorshift RNGs". Journal of Statistical Software. 8 (14).
type Source struct {
	state uint32
}

// New create a new source.
func New(state uint32) *Source {
	if state == 0 {
		state = 1
	}
	return &Source{
		state: state,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := int64(s.Uint32()&0x7FFFFFFF) << 32
	x += int64(s.Uint32())
	return x
}

// Uint32 returns pseudo-random uint32 values in the range [0, 1<<32).
func (s *Source) Uint32() uint32 {
	x := s.state
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	s.state = x
	return x
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	// use low 32bit for seed
	s.state = uint32(seed)
	if s.state != 0 {
		return
	}
	// state must be not 0, so fall back to high 32bit
	s.state = uint32(seed >> 32)
	if s.state != 0 {
		return
	}
	s.state = 1
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := uint64(s.Uint32()) << 32
	x += uint64(s.Uint32())
	return x
}
