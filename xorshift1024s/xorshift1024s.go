package xorshift1024s

import (
	"math/rand"

	"github.com/shogo82148/randsrc/xorshift64"
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using xorshift1024* algorithm.
//
// http://vigna.di.unimi.it/ftp/papers/xorshift.pdf
// https://arxiv.org/abs/1402.6246
// https://doi.org/10.1145%2F2845077
// Vigna, Sebastiano (July 2016). "An experimental exploration of Marsaglia's xorshift generators, scrambled" (PDF).
// ACM Transactions on Mathematical Software. 42 (4): 30. arXiv:1402.6246. doi:10.1145/2845077.
// Proposes xorshift* generators, adding a final multiplication by a constant.
type Source struct {
	state [16]uint64
	index int
}

// New create a new source.
func New(state [16]uint64) *Source {
	var zero [16]uint64
	if state == zero {
		state[0] = 1
	}
	return &Source{
		state: state,
		index: 0,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() & 0x7FFFFFFFFFFFFFFF
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	src := xorshift64.New(uint64(seed))
	for i := range s.state {
		s.state[i] = src.Uint64()
	}
	s.index = 0
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	index := s.index
	u := s.state[index]
	index = (index + 1) % 16
	t := s.state[index]
	t ^= t << 31
	t ^= t >> 11
	t ^= u ^ (u >> 30)
	s.state[index] = t
	s.index = index
	return t * 1181783497276652981
}
