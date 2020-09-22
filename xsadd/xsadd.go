package xsadd

import "math/rand"

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using XORSHIFT-ADD (XSadd) algorithm.
//
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/XSADD/index.html
// Go port of https://github.com/MersenneTwister-Lab/XSadd
type Source struct {
	state [4]uint32
}

// New create a new source.
func New(state [4]uint32) *Source {
	var zero [4]uint32
	if state == zero {
		state = [...]uint32{'X', 'S', 'A', 'D'}
	}
	return &Source{
		state: state,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := int64(s.Uint32()>>1) << 32
	x += int64(s.Uint32())
	return x
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	state := [4]uint32{uint32(seed)}
	for i := 1; i < 8; i++ {
		j := (i - 1) & 3
		state[i&3] ^= uint32(i) + 1812433253*(state[j]^(state[j]>>30))
	}
	var zero [4]uint32
	if state == zero {
		state = [...]uint32{'X', 'S', 'A', 'D'}
	}
	s.state = state
	for i := 0; i < 8; i++ {
		s.Uint32()
	}
}

func iniFunc1(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1664525
}

func iniFunc2(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1566083941
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint32) {
	const lag = 1
	const mid = 1
	const size = 4
	var state [4]uint32
	count := 8
	if len(initKey)+1 > 8 {
		count = len(initKey) + 1
	}

	r := iniFunc1(state[0] ^ state[mid%size] ^ state[(size-1)%size])
	state[mid%size] += r
	r += uint32(len(initKey))
	state[(mid+lag)%size] += r
	state[0] = r
	count--

	var i, j int
	for i, j = 1, 0; j < count && j < len(initKey); j++ {
		r := iniFunc1(state[i%size] ^ state[(i+mid)%size] ^ state[(i+size-1)%size])
		state[(i+mid)%size] += r
		r += initKey[j] + uint32(i)
		state[(i+mid+lag)%size] += r
		state[i%size] = r
		i = (i + 1) % size
	}
	for ; j < count; j++ {
		r := iniFunc1(state[i%size] ^ state[(i+mid)%size] ^ state[(i+size-1)%size])
		state[(i+mid)%size] += r
		r += uint32(i)
		state[(i+mid+lag)%size] += r
		state[i%size] = r
		i = (i + 1) % size
	}
	for j = 0; j < size; j++ {
		r := iniFunc2(state[i%size] + state[(i+mid)%size] + state[(i+size-1)%size])
		state[(i+mid)%size] ^= r
		r -= uint32(i)
		state[(i+mid+lag)%size] ^= r
		state[i%size] = r
		i = (i + 1) % size
	}

	var zero [4]uint32
	if state == zero {
		state = [...]uint32{'X', 'S', 'A', 'D'}
	}
	s.state = state
	for i := 0; i < 8; i++ {
		s.Uint32()
	}
}

// Uint32 returns pseudo-random uint32 values in the range [0, 1<<32).
func (s *Source) Uint32() uint32 {
	t := s.state[0]
	t ^= t << 15
	t ^= t >> 18
	t ^= s.state[3] << 11
	s.state = [...]uint32{s.state[1], s.state[2], s.state[3], t}
	return s.state[3] + s.state[2]
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := uint64(s.Uint32()) << 32
	x += uint64(s.Uint32())
	return x
}
