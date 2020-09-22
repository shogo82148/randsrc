package xoshiro256p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256p"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256p.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256p.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256p.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 1
	// 1
	// 35184372088832
	// 70368811286529
	// 4611721202933825664
	// 4512395720392960
	// 4620746020094115969
	// 1152930317980483617
	// 7501890900650557505
	// 1164265194092765472
}
