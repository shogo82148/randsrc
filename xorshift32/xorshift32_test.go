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
		fmt.Println(r.Int63())
	}
	//Output:
	// 580613040243456
	// 5685324361786641575
	// 5151145904873909736
	// 1358144856227876441
	// 4306488609506595258
	// 5706062264076749723
	// 6511120689350704215
	// 187840544559483608
	// 1914254710601805188
	// 689360569037983500
}
