package xoshiro256ss_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256ss"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256ss.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256ss.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256ss.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 0
	// 2880
	// 2880
	// 377490240
	// 101330991993323520
	// 49671674268480
	// 101380470016942080
	// 1008855794931728384
	// 6586539387786734400
	// 6510565882661532660
}
