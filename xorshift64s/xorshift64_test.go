package xorshift64s_test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/xorshift64s"
)

func BenchmarkInt63(b *testing.B) {
	src := xorshift64s.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := xorshift64s.New(1)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource() {
	src := xorshift64s.New(1)
	r := rand.New(src)
	for i := 0; i < 10; i++ {
		fmt.Println(r.Int())
	}
	//Output:
	// 5180492295206395165
	// 3156925108060775709
	// 4166126042076094295
	// 5599127315341312413
	// 1036278371763004928
	// 5217222029704669913
	// 5787885115471196545
	// 3202495810276243853
	// 6247250396617125944
	// 4610193123267394197
}
