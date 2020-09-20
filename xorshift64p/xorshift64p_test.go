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
	// 8388677
	// 33554692
	// 70368777736387
	// 211106267148357
	// 281509366091972
	// 360777324180299
	// 288538377073858266
	// 865509272901433454
	// 1155350481850751274
	// 1191592583033944276
}
