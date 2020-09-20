package splitmix64_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/splitmix64"
)

func BenchmarkInt63(b *testing.B) {
	src := splitmix64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := splitmix64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := splitmix64.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 1227844342346046657
	// 4533873174211652711
	// 8688467253428114782
	// 8196980753821780235
	// 8195237237126968761
	// 4849545566009754240
	// 6960854651289091237
	// 425514363213284725
	// 5266705631892356520
	// 5423280143191861142
}
