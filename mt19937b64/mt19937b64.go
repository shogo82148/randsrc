// Package mt19937b64 is an implementation of Mersenne Twister 64bit algorithm.
//
// References:
// T. Nishimura, ``Tables of 64-bit Mersenne Twisters''
//   ACM Transactions on Modeling and
//   Computer Simulation 10. (2000) 348--357.
// M. Matsumoto and T. Nishimura,
//   ``Mersenne Twister: a 623-dimensionally equidistributed
// 	uniform pseudorandom number generator''
//   ACM Transactions on Modeling and
//   Computer Simulation 8. (Jan. 1998) 3--30.
package mt19937b64

import (
	"math/rand"
)

// Period parameters
const (
	_NN      = 312
	_MM      = 156
	_MatrixA = 0xB5026F5AA96619E9 // constant vector a
	_UM      = 0xFFFFFFFF80000000 // Most significant 33 bits
	_LM      = 0x7FFFFFFF         // Least significant 31 bits
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using Mersenne Twister 64bit algorithm.
//
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html
type Source struct {
	mt  [_NN]uint64
	mti int
}

// New create a new source.
func New(mt [312]uint64) *Source {
	return &Source{
		mt:  mt,
		mti: _NN,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := s.Uint64() >> 1
	return int64(x)
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	state := uint64(seed)
	for i := range s.mt {
		s.mt[i] = state
		state = 6364136223846793005*(state^(state>>62)) + uint64(i+1)
	}
	s.mti = _NN
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint64) {
	s.Seed(19650218)
	i, j, k := 1, 0, len(initKey)
	if k < _NN {
		k = _NN
	}
	for ; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ ((s.mt[i-1] ^ (s.mt[i-1] >> 62)) * 3935559000370003845)) + initKey[j] + uint64(j)
		i++
		j++
		if i >= _NN {
			s.mt[0] = s.mt[_NN-1]
			i = 1
		}
		if j >= len(initKey) {
			j = 0
		}
	}
	for k = _NN - 1; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ ((s.mt[i-1] ^ (s.mt[i-1] >> 62)) * 2862933555777941757)) - uint64(i)
		i++
		if i >= _NN {
			s.mt[0] = s.mt[_NN-1]
			i = 1
		}
	}
	s.mt[0] = 1 << 63 // MSB is 1; assuring non-zero initial array
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	var x uint64

	mt := s.mt[:]
	if s.mti >= _NN {
		// generate NN words at one time
		var mag01 = [2]uint64{0, _MatrixA}
		var i int
		for i = 0; i < _NN-_MM; i++ {
			x = (mt[i] & _UM) | (mt[i+1] & _LM)
			mt[i] = mt[i+_MM] ^ (x >> 1) ^ mag01[x&1]
		}
		for ; i < _NN-1; i++ {
			x = (mt[i] & _UM) | (mt[i+1] & _LM)
			mt[i] = mt[i+(_MM-_NN)] ^ (x >> 1) ^ mag01[x&1]
		}
		x = (mt[_NN-1] & _UM) | (mt[0] & _LM)
		mt[_NN-1] = mt[_MM-1] ^ (x >> 1) ^ mag01[x&1]

		s.mti = 0
	}

	x = mt[s.mti]
	s.mti++

	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)

	return x
}
