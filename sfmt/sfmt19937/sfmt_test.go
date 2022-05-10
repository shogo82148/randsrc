package sfmt19937

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestSource_Uint64(t *testing.T) {
	src := New(4321)

	f, err := os.Open(filepath.Join("../", "testdata", fmt.Sprintf("m%d.txt", mexp)))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		got := src.Uint64()
		var want uint64
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt(%d) mismatch: want %016x, got %016x", i, want, got)
		}
	}
}

func TestSource_SeedBySlice(t *testing.T) {
	src := New(0)
	src.SeedBySlice([]uint32{5, 4, 3, 2, 1})

	f, err := os.Open(filepath.Join("../", "testdata", fmt.Sprintf("seedBySlice_m%d.txt", mexp)))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		got := src.Uint64()
		var want uint64
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt(%d) mismatch: want 0x%016x, got 0x%016x", i, want, got)
		}
	}
}

func BenchmarkInt63(b *testing.B) {
	src := New(4321)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := New(4321)
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}
