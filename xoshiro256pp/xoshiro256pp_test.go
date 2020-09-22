package xoshiro256pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256pp.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256pp.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256pp.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 4194304
	// 4194304
	// 8
	// 299067166949392
	// 563500280774664
	// 2305843010289534018
	// 110695128998414348
	// 2378323163791769602
	// 706296316889343765
	// 7326126286536182355
}
