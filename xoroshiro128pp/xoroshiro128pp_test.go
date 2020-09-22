package xoroshiro128pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro128pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro128pp.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro128pp.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoroshiro128pp.New([2]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 65536
	// 299204602822658
	// 289675143695892772
	// 5082613694791532834
	// 271823060665434799
	// 3628234659518064874
	// 7563027819649276380
	// 2173227447226003975
	// 8233774529164765107
	// 2037889483630097538
}
