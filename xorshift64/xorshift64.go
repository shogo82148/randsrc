package xorshift64

import "math/rand"

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorshift64 algorithm.
//
// Marsaglia, George (July 2003). "Xorshift RNGs". Journal of Statistical Software. 8 (14).
type Source struct {
	state uint64
}

// New create a new source.
func New(state uint64) *Source {
	if state == 0 {
		state = 1
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
	s.state = uint64(seed)
	if s.state != 0 {
		return
	}
	s.state = 1
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := s.state
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	s.state = x
	return x
}
