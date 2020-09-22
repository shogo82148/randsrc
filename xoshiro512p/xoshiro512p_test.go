package xoshiro512p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro512p"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro512p.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro512p.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro512p.New([8]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 0
	// 1
	// 0
	// 1025
	// 2048
	// 1024
	// 3147777
	// 2199028498432
	// 4611688217453789185
	// 4611688221751377920
}
