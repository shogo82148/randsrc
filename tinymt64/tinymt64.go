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

func iniFunc1(x uint64) uint64 {
	return (x ^ (x >> 59)) * 2173292883993
}

func iniFunc2(x uint64) uint64 {
	return (x ^ (x >> 59)) * 58885565329898161
}

// SeedBySlice initializes the state by initKey.
func (s *Source) SeedBySlice(initKey []uint64) {
	const lag = 1
	const mid = 1
	const size = 4
	var state [4]uint64
	count := 8
	if len(initKey)+1 > 8 {
		count = len(initKey) + 1
	}

	state[0] = 0
	state[1] = uint64(s.mat1)
	state[2] = uint64(s.mat2)
	state[3] = s.tmat

	r := iniFunc1(state[0] ^ state[mid%size] ^ state[(size-1)%size])
	state[mid%size] += r
	r += uint64(len(initKey))
	state[(mid+lag)%size] += r
	state[0] = r
	count--

	var i, j int
	for i, j = 1, 0; j < count && j < len(initKey); j++ {
		r := iniFunc1(state[i%size] ^ state[(i+mid)%size] ^ state[(i+size-1)%size])
		state[(i+mid)%size] += r
		r += initKey[j] + uint64(i)
		state[(i+mid+lag)%size] += r
		state[i%size] = r
		i = (i + 1) % size
	}
	for ; j < count; j++ {
		r := iniFunc1(state[i%size] ^ state[(i+mid)%size] ^ state[(i+size-1)%size])
		state[(i+mid)%size] += r
		r += uint64(i)
		state[(i+mid+lag)%size] += r
		state[i%size] = r
		i = (i + 1) % size
	}
	for j = 0; j < size; j++ {
		r := iniFunc2(state[i%size] + state[(i+mid)%size] + state[(i+size-1)%size])
		state[(i+mid)%size] ^= r
		r -= uint64(i)
		state[(i+mid+lag)%size] ^= r
		state[i%size] = r
		i = (i + 1) % size
	}

	s.status[0] = state[0] ^ state[1]
	s.status[1] = state[2] ^ state[3]

	var zero [2]uint64
	s.status[0] &= mask
	if s.status == zero {
		s.status = [2]uint64{'T', 'M'}
	}
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
