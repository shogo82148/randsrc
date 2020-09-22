package xoshiro256pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256pp.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256pp.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256pp.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 8388609
	// 8388609
	// 16
	// 598134333898785
	// 1127000561549328
	// 4611686020579068036
	// 221390257996828697
	// 4756646327583539205
	// 1412592633778687530
	// 5428880536217588902
}
