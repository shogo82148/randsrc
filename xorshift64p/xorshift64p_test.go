package xorshift64p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift64p"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift64p.New(1, 2)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift64p.New(1, 2)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift64p.New(1, 2)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 4194338
	// 16777346
	// 35184388868193
	// 105553133574178
	// 140754683045986
	// 180388662090149
	// 144269188536929133
	// 432754636450716727
	// 577675240925375637
	// 595796291516972138
}
