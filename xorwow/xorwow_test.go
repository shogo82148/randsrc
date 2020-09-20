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
		fmt.Println(r.Int())
	}
	//Output:
	// 1556779616637260
	// 4702843161797285
	// 16221485918290218
	// 2165720138257341632
	// 8850039960330671171
	// 7330236657558040176
	// 6368334201278476546
	// 63177036399988055
	// 7029260376017917067
	// 7968997030677343245
}
