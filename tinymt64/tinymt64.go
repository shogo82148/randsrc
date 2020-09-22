/*
* Copyright (C) 2020 Ichinose, Shogo
* Copyright (C) 2011 Mutsuo Saito, Makoto Matsumoto,
* Hiroshima University and The University of Tokyo.
* All rights reserved.
*
* The 3-clause BSD License is applied to this software, see
* LICENSE.txt
 */

package tinymt64

import (
	"math/rand"
)

const (
	sh0  = 12
	sh1  = 11
	sh8  = 8
	mask = 0x7fffffffffffffff
)

var _ rand.Source = (*Source)(nil)
var _ rand.Source64 = (*Source)(nil)

// Source is a random source using Tiny Mersenne Twister(TinyMT) algorithm.
//
// Go port of https://github.com/MersenneTwister-Lab/TinyMT
type Source struct {
	status     [2]uint64
	mat1, mat2 uint32
	tmat       uint64
}

// New creates a new source.
// mat1, mat2, and tmat are a parameter set that needs to be well chosen.
// the precalculated parameter sets are available at https://github.com/jj1bdx/tinymtdc-longbatch
func New(mat1, mat2 uint32, tmat uint64, status [2]uint64) *Source {
	var zero [2]uint64
	status[0] &= mask
	if status == zero {
		status = [2]uint64{'T', 'M'}
	}
	return &Source{
		status: status,
		mat1:   mat1,
		mat2:   mat2,
		tmat:   tmat,
	}
}

// Int63 implements math/rand.Source.
func (s *Source) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *Source) next() {
	s.status[0] &= mask
	x := s.status[0] ^ s.status[1]
	x ^= x << sh0
	x ^= x >> 32
	x ^= x << 32
	x ^= x << sh1
	s.status[0] = s.status[1]
	s.status[1] = x
	if x%2 != 0 {
		s.status[0] ^= uint64(s.mat1)
		s.status[1] ^= uint64(s.mat2) << 32
	}
}

// Seed implements math/rand.Source.
func (s *Source) Seed(seed int64) {
	s.status[0] = uint64(seed) ^ (uint64(s.mat1) << 32)
	s.status[1] = uint64(s.mat2) ^ s.tmat
	for i := 1; i < 8; i++ {
		s.status[i&1] ^= uint64(i) + 6364136223846793005*(s.status[(i-1)&1]^(s.status[(i-1)&1]>>62))
	}
	var zero [2]uint64
	s.status[0] &= mask
	if s.status == zero {
		s.status = [2]uint64{'T', 'M'}
	}
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint32) {
}

// Uint64 implements math/rand.Source64
func (s *Source) Uint64() uint64 {
	s.next()
	x := s.status[0] + s.status[1]
	x ^= s.status[0] >> sh8
	if x%2 != 0 {
		x ^= s.tmat
	}
	return x
}
