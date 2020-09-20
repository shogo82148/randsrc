package xorshift1024s_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift1024s"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift1024s.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift1024s.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift1024s.New([16]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
	// 1181783497276652981
}
