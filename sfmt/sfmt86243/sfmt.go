package sfmt86243

// Code generated ../generate.sh; DO NOT EDIT.

import (
	"math/bits"
	"math/rand"
)

// w128t is unsigned 128-bit integer type.
type w128t [2]uint64

const maskR = uint64((0xFFFFFFFF >> sr1) | uint64(0xFFFFFFFF>>sr1)<<32)
const maskL = uint64((0xFFFFFFFF & (0xFFFFFFFF << sl1)) | (0xFFFFFFFF&(0xFFFFFFFF<<sl1))<<32)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using Mersenne Twister algorithm.
type Source struct {
	state [(mexp >> 7) + 1]w128t
	idx   int
}

// New create a new source.
func New(seed int64) *Source {
	s := &Source{}
	s.Seed(seed)
	return s
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() >> 1
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	x := uint32(seed)
	for i := range s.state {
		a := x
		x = 1812433253*(x^(x>>30)) + uint32(i)*4 + 1
		b := x
		x = 1812433253*(x^(x>>30)) + uint32(i)*4 + 2
		c := x
		x = 1812433253*(x^(x>>30)) + uint32(i)*4 + 3
		d := x
		x = 1812433253*(x^(x>>30)) + uint32(i)*4 + 4
		s.state[i] = w128t{
			uint64(a) | (uint64(b) << 32),
			uint64(c) | (uint64(d) << 32),
		}
	}
	s.periodCertification()
	s.idx = len(s.state) * 2
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint32) {
	size := len(s.state) * 4
	state := make([]uint32, size)
	var lag int
	switch {
	case size >= 623:
		lag = 11
	case size >= 68:
		lag = 7
	case size >= 39:
		lag = 5
	default:
		lag = 3
	}
	mid := (size - lag) / 2

	for i := range state {
		state[i] = 0x8b8b8b8b
	}

	count := size
	if len(initKey)+1 > count {
		count = len(initKey) + 1
	}

	r := func1(state[0] ^ state[mid] ^ state[size-1])
	state[mid] += r
	r += uint32(len(initKey))
	state[mid+lag] += r
	state[0] = r

	count--
	i, j := 1, 0
	for ; j < count && j < len(initKey); j++ {
		r = func1(state[i] ^ state[(i+mid)%size] ^ state[((i+size-1)%size)])
		state[(i+mid)%size] += r
		r += initKey[j] + uint32(i)
		state[(i+mid+lag)%size] += r
		state[i] = r
		i = (i + 1) % size
	}
	for ; j < count; j++ {
		r = func1(state[i] ^ state[(i+mid)%size] ^ state[((i+size-1)%size)])
		state[(i+mid)%size] += r
		r += uint32(i)
		state[(i+mid+lag)%size] += r
		state[i] = r
		i = (i + 1) % size
	}
	for j := 0; j < size; j++ {
		r = func2(state[i] + state[(i+mid)%size] + state[((i+size-1)%size)])
		state[(i+mid)%size] ^= r
		r -= uint32(i)
		state[(i+mid+lag)%size] ^= r
		state[i] = r
		i = (i + 1) % size
	}

	for i := range s.state {
		s.state[i][0] = (uint64(state[i*4+1]) << 32) | uint64(state[i*4+0])
		s.state[i][1] = (uint64(state[i*4+3]) << 32) | uint64(state[i*4+2])
	}

	s.periodCertification()
	s.idx = len(s.state) * 2
}

// This function certificate the period of 2^{MEXP}
func (s *Source) periodCertification() {
	inner := (s.state[0][0] & parity1) ^ (s.state[0][1] & parity2)
	if bits.OnesCount64(inner)%2 != 0 {
		// check OK
		return
	}

	// check NG, and modification
	if parity1 != 0 {
		s.state[0][0] ^= 1 << bits.TrailingZeros64(parity1)
	} else if parity2 != 0 {
		s.state[0][1] ^= 1 << bits.TrailingZeros64(parity2)
	} else {
		panic("sfmt: period certification failed")
	}
}
func func1(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1664525
}

func func2(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1566083941
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	if s.idx >= len(s.state)*2 {
		var i, j int
		n := len(s.state)
		r1 := s.state[n-2]
		r2 := s.state[n-1]
		j = int(pos1)
		for j < n {
			s.state[i] = s.dorecursion(
				s.state[i], s.state[j], r1, r2,
			)
			r1 = r2
			r2 = s.state[i]
			i++
			j++
		}

		j = 0
		for i < n {
			s.state[i] = s.dorecursion(
				s.state[i], s.state[j], r1, r2,
			)
			r1 = r2
			r2 = s.state[i]
			i++
			j++
		}
		s.idx = 0
	}

	ret := s.state[s.idx/2][s.idx%2]
	s.idx++
	return ret
}

func (s *Source) dorecursion(a, b, c, d w128t) w128t {
	// inlining by hand, because it is too deep for Go to optimize function calls.
	// x := a.lshift(s.param.SL2 * 8)
	x := w128t{
		a[0] << (sl2 * 8),
		(a[1] << (sl2 * 8)) | (a[0] >> (64 - sl2*8)),
	}

	// y := b.rshift(s.s.param.SR2)
	y := w128t{
		(c[0] >> (sr2 * 8)) | (c[1] << (64 - sr2*8)),
		c[1] >> (sr2 * 8),
	}

	var r w128t
	r[0] = a[0] ^ x[0] ^ ((b[0] >> sr1) & msk1) ^ y[0] ^ ((d[0] << sl1) & maskL)
	r[1] = a[1] ^ x[1] ^ ((b[1] >> sr1) & msk2) ^ y[1] ^ ((d[1] << sl1) & maskL)
	return r
}
