package xoroshiro1024ss_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro1024ss"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro1024ss.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro1024ss.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	var src xoroshiro1024ss.Source
	src.Seed(1)
	r := rand.New(&src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 6483309580052039778
	// 9000725922818581354
	// 4695528695355784254
	// 8897572651153514942
	// 2505966309961638356
	// 4873959627154931996
	// 7996257189429601695
	// 4888591260735349491
	// 3877953997424646574
	// 7921896405461697353
}
