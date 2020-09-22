package xoroshiro128ss_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro128ss"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro128ss.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro128ss.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoroshiro128ss.New([2]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int63())
	}
	//Output:
	// 2880
	// 48507128640
	// 8305045906563009344
	// 4611698629741136713
	// 7912389908774368352
	// 6226882119986805214
	// 4157124678997014974
	// 7707811687779445541
	// 8305244408713002896
	// 8529798522661292347
}
