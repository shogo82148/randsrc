package xoroshiro128pp_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xoroshiro128pp"
)

func BenchmarkInt63(b *testing.B) {
	src := xoroshiro128pp.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xoroshiro128pp.New([2]uint64{1})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xoroshiro128pp.New([2]uint64{1})
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 131073
	// 598409205645317
	// 579350287391785545
	// 941855352728289861
	// 543646121330869598
	// 7256469319036129749
	// 5902683602443776953
	// 4346454894452007950
	// 7244177021474754406
	// 4075778967260195077
}
