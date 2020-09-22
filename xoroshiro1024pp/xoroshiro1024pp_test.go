package xoroshiro1024pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro1024pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro1024pp.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro1024pp.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	var src xoroshiro1024pp.Source
	src.Seed(1)
	r := rand.New(&src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 3170740057973009359
	// 4099357889825430042
	// 5638365265662465543
	// 4684694968524113019
	// 5286072165563643916
	// 8404018247976732837
	// 7570187217392227886
	// 5076094242191494162
	// 5427123157498930871
	// 1232801326997470226
}
