package xoroshiro128p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro128p"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro128p.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro128p.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoroshiro128p.New([2]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 0
	// 68727898112
	// 1162069580936380928
	// 36099245263065744
	// 1306769709744035040
	// 2568221860186161298
	// 3046776487690310554
	// 8213190904191315488
	// 3214698032120034443
	// 292334754365524544
}
