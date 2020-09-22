package xoroshiro128p_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro128p"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro128p.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro128p.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoroshiro128p.New([2]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 1
	// 137455796225
	// 2324139161872761857
	// 72198490526131489
	// 2613539419488070081
	// 5136443720372322597
	// 6093552975380621109
	// 7203009771527855168
	// 6429396064240068886
	// 584669508731049089
}
