package xorshift32_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift32"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift32.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift32.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift32.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 580610926577153
	// 5685324359792957775
	// 5151145903099173840
	// 1358144856445754546
	// 4306488608817161076
	// 5706062262676789047
	// 6511120689391728814
	// 187840544104551857
	// 1914254712263400200
	// 689360569032449561
}
