package xorshift128_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift128"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift128.New(1, 2)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift128.New(1, 2)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift128.New(1, 2)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 4114
	// 17669495459858
	// 17669495461915
	// 26504243189787
	// 26504251578505
	// 36055773708556299
	// 26435536296144
	// 54070477160716307
	// 26469887904970
	// 19154360147938441
}
