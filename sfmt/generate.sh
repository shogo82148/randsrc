#!/bin/bash

set -ue

sfmt() {
    local MEXP=$1
	local Pos1=$2
	local SL1=$3
	local SL2=$4
	local SR1=$5
	local SR2=$6
	local MSK1=$7
	local MSK2=$8
	local Parity1=$9
	local Parity2=${10}

    cat <<EOF | gofmt > "param${MEXP}.go"
package sfmt

// Code generated generate.sh; DO NOT EDIT.

// Param${MEXP} xxx
var Param${MEXP} = &Parameter{
	MExp:    ${MEXP},
	Pos1:    ${Pos1},
	SL1:     ${SL1},
	SL2:     ${SL2},
	SR1:     ${SR1},
	SR2:     ${SR2},
	MSK1:    ${MSK1},
	MSK2:    ${MSK2},
	Parity1: ${Parity1},
	Parity2: ${Parity2},
}
EOF
    mkdir -p "sfmt${MEXP}"

    cat <<EOF | gofmt > "sfmt${MEXP}/param.go"
package sfmt${MEXP}

// Code generated ../generate.sh; DO NOT EDIT.

const (
	// Mersenne Exponent. The period of the sequence
	// is a multiple of 2^MEXP-1.
	mexp = ${MEXP}

	// the pick up position of the array.
	pos1 = ${Pos1}

	// the parameter of shift left as four 32-bit registers.
	sl1 = ${SL1}

	// the parameter of shift left as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SL2 * 8) bits.
	sl2 = ${SL2}

	// the parameter of shift right as four 32-bit registers.
	sr1 = ${SR1}

	// the parameter of shift right as one 128-bit register.
	// The 128-bit integer is shifted by (SFMT_SR2 * 8) bits.
	sr2 = ${SR2}

	// A bitmask, used in the recursion.
	// These parameters are introduced to break symmetry of SIMD.
	msk1 = ${MSK1} & maskR
	msk2 = ${MSK2} & maskR

	// These definitions are part of a 128-bit period certification vector.
	parity1 = ${Parity1}
	parity2 = ${Parity2}
)
EOF
    cat <<EOF | gofmt > "sfmt${MEXP}/sfmt.go"
package sfmt${MEXP}

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
EOF

    cat <<EOF | gofmt >  "sfmt${MEXP}/sfmt_test.go"
package sfmt${MEXP}

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSource_Uint64(t *testing.T) {
	src := New(4321)

	f, err := os.Open(filepath.Join("../", "testdata", fmt.Sprintf("m%d.txt", mexp)))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		got := src.Uint64()
		var want uint64
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt(%d) mismatch: want %016x, got %016x", i, want, got)
		}
	}
}

func TestSource_SeedBySlice(t *testing.T) {
	src := New(0)
	src.SeedBySlice([]uint32{5, 4, 3, 2, 1})

	f, err := os.Open(filepath.Join("../", "testdata", fmt.Sprintf("seedBySlice_m%d.txt", mexp)))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		got := src.Uint64()
		var want uint64
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt(%d) mismatch: want 0x%016x, got 0x%016x", i, want, got)
		}
	}
}

func BenchmarkInt63(b *testing.B) {
	src := New(4321)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := New(4321)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}
EOF
}

sfmt    607   2 15 3 13 3 0xef7f3f7dfdff37ff 0x7ff7fb2fff777b7d 0x0000000000000001 0x5986f05400000000
sfmt   1279   7 14 3  5 1 0x7fefcffff7fefffd 0xb5ffff7faff3ef3f 0x0000000000000001 0x2000000000000000
sfmt   2281  12 19 1  5 1 0xfdfffffebff7ffbf 0xf2f7cbbff7ffef7f 0x0000000000000001 0x41dfa60000000000
sfmt   4253  17 20 1  7 1 0x9fffff5f9f7bffff 0xfffff7bb3efffffb 0xaf5390a3a8000001 0x6c11486db740b3f8
sfmt  11213  68 14 3  7 3 0xffffffefeffff7fb 0x7fffdbfddfdfbfff 0x0000000000000001 0xd0c7afa3e8148000
sfmt  19937 122 18 1 11 1 0xddfecb7fdfffffef 0xbffffff6bffaffff 0x0000000000000001 0x13c9e68400000000
sfmt  44497 330  5 3  9 3 0xdfbebfffeffffffb 0x9ffd7bffbfbf7bef 0x0000000000000001 0xecc1327aa3ac4000
sfmt  86243 366  6 7 19 1 0xbff7ff3ffdbffbff 0xbf9ff3fffd77efff 0x0000000000000001 0xe9528d8500000000
sfmt 132049 110 19 1 21 1 0xfb6ebf95ffffbb5f 0xcff77ffffffefffa 0x0000000000000001 0xc7e91c7dcb520000
sfmt 216091 627 11 3 10 1 0xbfffffffbff7bff7 0xffddfbfbbffffa7f 0x89e80709f8000001 0x0c64b1e43bd2b64b
