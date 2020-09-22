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
	// 778389808318630
	// 2351421580898642
	// 8110742959145109
	// 1082860069128670816
	// 9036705998592723489
	// 3665118328779020088
	// 7795853119066626177
	// 4643274536627381931
	// 8126316206436346437
	// 8596184533766059526
}
