package splitmix64

import (
	"math/rand"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using splitmix64 algorithm.
//
// Guy L. Steele, Doug Lea, and Christine H. Flood. 2014.
// Fast splittable pseudorandom number generators. SIGPLAN Not. 49, 10 (October 2014), 453–472.
// DOI:https://doi.org/10.1145/2714064.2660195
//
// http://docs.oracle.com/javase/8/docs/api/java/util/SplittableRandom.html
// http://prng.di.unimi.it/splitmix64.c
type Source struct {
	state uint64
}

// New create a new source.
func New(state uint64) *Source {
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
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	s.state += 0x9e3779b97f4a7c15
	z := s.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
