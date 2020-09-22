package xorwow_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorwow"
)

func BenchmarkInt63(b *testing.B) {
	src := xorwow.New(1, 2, 3, 4, 5)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorwow.New(1, 2, 3, 4, 5)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorwow.New(1, 2, 3, 4, 5)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 778389808681292
	// 2351421581681317
	// 8110740827504938
	// 1082860068623162560
	// 9036705999293355075
	// 3665118326877269616
	// 7795853117105156354
	// 4643274536732937559
	// 8126316205918573707
	// 8596184533612260365
}
