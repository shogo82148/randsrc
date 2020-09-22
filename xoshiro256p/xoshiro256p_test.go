package xoshiro256p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256p"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256p.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256p.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256p.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 0
	// 0
	// 17592186044416
	// 35184405643264
	// 2305860601466912832
	// 4613942216287584384
	// 2310373010047057984
	// 5188151177417629712
	// 3750945450325278752
	// 5193818615473770640
}
