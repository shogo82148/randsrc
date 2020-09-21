package mt19937b64_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/mt19937"
)

func BenchmarkInt63(b *testing.B) {
	src := &mt19937.Source{}
	src.Seed(5489)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := &mt19937.Source{}
	src.Seed(5489)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	var src mt19937.Source
	src.SeedBySlice([]uint32{0x123, 0x234, 0x345, 0x456})
	for i := 0; i < 10; i++ {
		fmt.Println(src.Uint32())
	}
	//Output:
	// 1067595299
	// 955945823
	// 477289528
	// 4107218783
	// 4228976476
	// 3344332714
	// 3355579695
	// 227628506
	// 810200273
	// 2591290167
}
