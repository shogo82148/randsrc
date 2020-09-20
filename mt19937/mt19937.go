package mt19937

import (
	"math/rand"
)

// Period parameters
const (
	_N         = 624
	_M         = 397
	_MatrixA   = 0x9908b0df // constant vector a
	_UpperMask = 0x80000000 // most significant w-r bits
	_LowerMask = 0x7fffffff // least significant r bits
)

var mag01 = [2]uint32{0, _MatrixA}

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using Mersenne Twister algorithm.
//
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/emt19937ar.html
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/CODES/mt19937ar.c
type Source struct {
	mt  [_N]uint32
	mti int
}

// New create a new source.
func New(mt [624]uint32) *Source {
	return &Source{
		mt:  mt,
		mti: _N,
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
	var y uint32
	if s.mti >= len(s.mt) {
		// generate N words at one time
		var kk int
		for ; kk < _N-_M; kk++ {
			y = (s.mt[kk] & _UpperMask) | (s.mt[kk+1] & _LowerMask)
			s.mt[kk] = s.mt[kk+_M] ^ (y >> 1) ^ mag01[y%2]
		}
		for ; kk < _N-1; kk++ {
			y = (s.mt[kk] & _UpperMask) | (s.mt[kk+1] & _LowerMask)
			s.mt[kk] = s.mt[kk+(_M-_N)] ^ (y >> 1) ^ mag01[y%2]
		}
		y = (s.mt[_N-1] & _UpperMask) | (s.mt[0] & _LowerMask)
		s.mt[_N-1] = s.mt[_M-1] ^ (y >> 1) ^ mag01[y%2]
		s.mti = 0
	}

	y = s.mt[s.mti]
	s.mti++

	// Tempering
	y ^= y >> 11
	y ^= (y << 7) & 0x9d2c5680
	y ^= (y << 15) & 0xefc60000
	y ^= (y >> 18)
	return y
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	state := uint32(seed)
	for i := range s.mt {
		s.mt[i] = state
		state = 1812433253*(state^(state>>30)) + uint32(i+1)
	}
	s.mti = _N
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint32) {
	s.Seed(19650218)

	i, j, k := 1, 0, len(initKey)
	if k < _N {
		k = _N
	}
	for ; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ (s.mt[i-1]^(s.mt[i-1]>>30))*1664525) + initKey[j] + uint32(j)
		i++
		j++
		if i >= _N {
			s.mt[0] = s.mt[_N-1]
			i = 1
		}
		if j >= len(initKey) {
			j = 0
		}
	}
	for k = _N - 1; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ (s.mt[i-1]^(s.mt[i-1]>>30))*1566083941) - uint32(i)
		i++
		if i >= _N {
			s.mt[0] = s.mt[_N-1]
			i = 1
		}
	}
	s.mt[0] = 0x80000000
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := uint64(s.Uint32()) << 32
	x += uint64(s.Uint32())
	return x
}
