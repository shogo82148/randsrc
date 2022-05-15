package bbs

import (
	"runtime"
	"testing"
)

func BenchmarkInt63(b *testing.B) {
	s := New(123)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(s.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	s := New(123)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(s.Uint64())
	}
}
