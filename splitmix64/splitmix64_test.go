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
	// 5225608189600411232
	// 6878622605533214259
	// 8955919645141445295
	// 4098490376910890117
	// 4097618618563484380
	// 7036458801432265024
	// 8092113344071933522
	// 4824443200034030266
	// 2633352815946178260
	// 7323326090023318475
}
