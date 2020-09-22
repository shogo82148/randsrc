package xorshift64s_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift64s"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift64s.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift64s.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift64s.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 2590246147603197582
	// 6190148572457775758
	// 6694749039465435051
	// 2799563657670656206
	// 518139185881502464
	// 7220297033279722860
	// 7505628576162986176
	// 6212933923565509830
	// 3123625198308562972
	// 6916782580061085002
}
