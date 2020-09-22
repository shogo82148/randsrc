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
	// 2057
	// 8834747729929
	// 8834747730957
	// 13252121594893
	// 13252125789252
	// 18027886854278149
	// 13217768148072
	// 27035238580358153
	// 13234943952485
	// 9577180073969220
}
