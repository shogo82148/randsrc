// The sfmt package implements SIMD-oriented Fast Mersenne Twister (SFMT):
// twice faster than Mersenne Twister.
//
// http://www.math.sci.hiroshima-u.ac.jp/m-mat/MT/SFMT/index.html
package sfmt

import (
	"fmt"
	"math/bits"
	"math/rand"
)

//go:generate ./generate.sh

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Parameter is a parameter of Source.
type Parameter struct {
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	MExp uint

	// the pick up position of the array.
	Pos1 uint

	// the parameter of shift left as four 32-bit registers.
	SL1 uint

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	SL2 uint

	// the parameter of shift right as four 32-bit registers.
	SR1 uint

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	SR2 uint

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	MSK1, MSK2 uint64

	// These definitions are part of a 128-bit period certification vector.
	Parity1, Parity2 uint64
}

func (p *Parameter) String() string {
	return fmt.Sprintf(
		"SFMT-%d:%d-%d-%d-%d-%d:%08x-%08x-%08x-%08x",
		p.MExp, p.Pos1, p.SL1, p.SL2, p.SR1, p.SR2,
		p.MSK1&0xFFFFFFFF, p.MSK1>>32, p.MSK2&0xFFFFFFFF, p.MSK2>>32,
	)
}

// Source is a random source using Mersenne Twister algorithm.
type Source struct {
	param Parameter
	state []w128t
	maskL uint64
	idx   int
}

// New create a new source.
func New(param *Parameter) *Source {
	p := *param
	maskR := uint64(uint32(0xFFFFFFFF) >> p.SR1)
	maskR |= maskR << 32
	p.MSK1 &= maskR
	p.MSK2 &= maskR
	maskL := uint64(uint32(0xFFFFFFFF) << p.SL1)
	maskL |= maskL << 32
	return &Source{
		param: p,
		maskL: maskL,
		state: make([]w128t, param.MExp/128+1),
	}
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

func func1(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1664525
}

func func2(x uint32) uint32 {
	return (x ^ (x >> 27)) * 1566083941
}

// This function certificate the period of 2^{MEXP}
func (s *Source) periodCertification() {
	inner := (s.state[0][0] & s.param.Parity1) ^ (s.state[0][1] & s.param.Parity2)
	if bits.OnesCount64(inner)%2 != 0 {
		// check OK
		return
	}

	// check NG, and modification
	if s.param.Parity1 != 0 {
		s.state[0][0] ^= 1 << bits.TrailingZeros64(s.param.Parity1)
	} else if s.param.Parity2 != 0 {
		s.state[0][1] ^= 1 << bits.TrailingZeros64(s.param.Parity2)
	} else {
		panic("sfmt: period certification failed")
	}
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	if s.idx >= len(s.state)*2 {
		var i, j int
		n := len(s.state)
		r1 := s.state[n-2]
		r2 := s.state[n-1]
		j = int(s.param.Pos1)
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
		a[0] << (s.param.SL2 * 8),
		(a[1] << (s.param.SL2 * 8)) | (a[0] >> (64 - s.param.SL2*8)),
	}

	// y := b.rshift(s.s.param.SR2)
	y := w128t{
		(c[0] >> (s.param.SR2 * 8)) | (c[1] << (64 - s.param.SR2*8)),
		c[1] >> (s.param.SR2 * 8),
	}

	var r w128t
	r[0] = a[0] ^ x[0] ^ ((b[0] >> s.param.SR1) & s.param.MSK1) ^ y[0] ^ ((d[0] << s.param.SL1) & s.maskL)
	r[1] = a[1] ^ x[1] ^ ((b[1] >> s.param.SR1) & s.param.MSK2) ^ y[1] ^ ((d[1] << s.param.SL1) & s.maskL)
	return r
}

// w128t is unsigned 128-bit integer type.
type w128t [2]uint64

// lshift returns x << n
// func (x w128t) lshift(n uint) w128t {
// 	return w128t{
// 		x[0] << (n * 8),
// 		(x[1] << (n * 8)) | (x[0] >> (64 - n*8)),
// 	}
// }

// rshift returns x >> n
// func (x w128t) rshift(n uint) w128t {
// 	return w128t{
// 		(x[0] >> (n * 8)) | (x[1] << (64 - n*8)),
// 		x[1] >> (n * 8),
// 	}
// }
