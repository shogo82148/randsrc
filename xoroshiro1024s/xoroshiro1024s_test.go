package xoroshiro1024s_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro1024s"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro1024s.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro1024s.New([16]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	var src xoroshiro1024s.Source
	src.Seed(1)
	r := rand.New(&src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 5342038924001916882
	// 2683941550217555709
	// 6531581289886261224
	// 3106969537905429597
	// 6043697972756256704
	// 5421517191659664927
	// 3817051424625445335
	// 2737981803588876732
	// 1056458538821463057
	// 1993465107093549081
}
