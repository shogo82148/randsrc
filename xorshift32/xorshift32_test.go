package xorshift32_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift32"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift32.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift32.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift32.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	//1161226080486913
	//2147276686718507343
	//1078919772893043664
	//2716289712455752882
	//8612977219013190516
	//2188752491298723639
	//3798869341846632622
	//375681089118967217
	//3828509421203610376
	//1378721138075967001
}
