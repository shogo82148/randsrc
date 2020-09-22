/*
* Copyright (C) 2020 Ichinose, Shogo
* Copyright (C) 2011 Mutsuo Saito, Makoto Matsumoto,
* Hiroshima University and The University of Tokyo.
* All rights reserved.
*
* The 3-clause BSD License is applied to this software, see
* LICENSE.txt
 */

package tinymt32

import (
	"math/rand"
)

const (
	sh0  = 1
	sh1  = 10
	sh8  = 8
	mask = 0x7fffffff
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using Tiny Mersenne Twister(TinyMT) algorithm.
//
// Go port of https://github.com/MersenneTwister-Lab/TinyMT
type Source struct {
	status           [4]uint32
	mat1, mat2, tmat uint32
}

// New creates a new source.
// mat1, mat2, and tmat are a parameter set that needs to be well chosen.
// the precalculated parameter sets are available at https://github.com/jj1bdx/tinymtdc-longbatch
func New(mat1, mat2, tmat uint32, status [4]uint32) *Source {
	return &Source{
		status: status,
		mat1:   mat1,
		mat2:   mat2,
		tmat:   tmat,
	}
}

// NewRFC8682 creates new source that is defined by RFC8682.
// https://trac.tools.ietf.org/html/rfc8682
func NewRFC8682(seed uint32) *Source {
	// https://trac.tools.ietf.org/html/rfc8682#section-2.1
	src := New(0x8f7011ee, 0xfc78ff1f, 0x3793fdff, [4]uint32{})
	src.Seed(int64(seed))
	return src
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	x := int64(s.Uint32()) << 31
	x += int64(s.Uint32()) >> 1
	return x
}

// Uint32 returns pseudo-random uint32 values in the range [0, 1<<32).
func (s *Source) Uint32() uint32 {
	s.next()
	t0 := s.status[3]
	t1 := s.status[0] + (s.status[2] >> sh8)
	t0 ^= t1
	t0 ^= -(t1 & 1) & s.tmat
	return t0
}

func (s *Source) next() {
	y := s.status[3]
	x := (s.status[0] & mask) ^ s.status[1] ^ s.status[2]
	x ^= (x << sh0)
	y ^= (y >> sh0) ^ x
	s.status[0] = s.status[1]
	s.status[1] = s.status[2]
	s.status[2] = x ^ (y << sh1)
	s.status[3] = y

	s.status[1] ^= -(y & 1) & s.mat1
	s.status[2] ^= -(y & 1) & s.mat2
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	s.status[0] = uint32(seed)
	s.status[1] = s.mat1
	s.status[2] = s.mat2
	s.status[3] = s.tmat
	for i := 1; i < 8; i++ {
		s.status[i&3] ^= uint32(i) + 1812433253*(s.status[(i-1)&3]^(s.status[(i-1)&3]>>30))
	}
	var zero [4]uint32
	if s.status == zero {
		s.status = [4]uint32{'T', 'I', 'N', 'Y'}
	}
	for i := 0; i < 8; i++ {
		s.next()
	}
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint32) {
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	x := uint64(s.Uint32()) << 32
	x += uint64(s.Uint32())
	return x
}
