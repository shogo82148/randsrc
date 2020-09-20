package mt19937

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSourceSeed(t *testing.T) {
	var src Source
	src.Seed(19650218)

	f, err := os.Open(filepath.Join("testdata", "seed19650218.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i, got := range src.mt {
		var want uint32
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt[%d] mismatch: want %d, got %d", i, want, got)
		}
	}
}

func TestSourceSeedBySlice(t *testing.T) {
	var src Source
	src.SeedBySlice([]uint32{0x123, 0x234, 0x345, 0x456})

	f, err := os.Open(filepath.Join("testdata", "seedBySlice.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i, got := range src.mt {
		var want uint32
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt[%d] mismatch: want %d, got %d", i, want, got)
		}
	}
}

func TestUint32(t *testing.T) {
	var src Source
	src.SeedBySlice([]uint32{0x123, 0x234, 0x345, 0x456})

	f, err := os.Open(filepath.Join("testdata", "list1000.txt"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		got := src.Uint32()
		var want uint32
		fmt.Fscanf(f, "%d", &want)
		if want != got {
			t.Errorf("mt[%d] mismatch: want %d, got %d", i, want, got)
		}
	}
}
