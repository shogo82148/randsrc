package xoshiro512ss_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro512ss"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro512ss.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro512ss.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro512ss.New([8]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 0
	// 2880
	// 2880
	// 0
	// 5898240
	// 5901120
	// 0
	// 18119393280
	// 12666386031577920
	// 6039800928
}
