package crypto

import (
	"runtime"
	"testing"
)

func BenchmarkInt63(b *testing.B) {
	src := New()
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := New()
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}
