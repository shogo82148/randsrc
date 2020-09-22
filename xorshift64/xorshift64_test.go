package xorshift64_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift64"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift64.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift64.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 541134880
	// 576496499416926752
	// 5588758332216382228
	// 8839011916000968722
	// 4829565071999682866
	// 8887899500566907904
	// 5846517310170255160
	// 7139982085056146566
	// 4512327100996027519
	// 4166335364516418199
}
