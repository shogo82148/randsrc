package xoshiro512pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro512pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro512pp.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro512pp.New([8]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro512pp.New([8]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 65536
	// 131072
	// 65536
	// 134348800
	// 268502016
	// 134283264
	// 412585427968
	// 288231063349690368
	// 288232987493040128
	// 4900482299473264640
}
