package xoshiro256ss_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoshiro256ss"
)

func BenchmarkInt63(b *testing.B) {
	src := xoshiro256ss.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoshiro256ss.New([4]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoshiro256ss.New([4]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 0
	// 5760
	// 5760
	// 754980480
	// 202661983986647040
	// 99343348536960
	// 202760940033884160
	// 2017711589863456768
	// 3949706738718692992
	// 3797759728468289512
}
