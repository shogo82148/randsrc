package xorshift64_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift64"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift64.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 1082269761
	// 1152992998833853505
	// 1954144627577988649
	// 8454651795147161637
	// 435758107144589925
	// 8552426964279040001
	// 2469662583485734513
	// 5056592133257517325
	// 9024654201992055039
	// 8332670729032836398
}
