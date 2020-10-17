package sfmt

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSource_Uint64(t *testing.T) {
	params := []*Parameter{
		Param607,
		Param1279,
		Param2281,
		Param4253,
		Param11213,
		Param19937,
		Param44497,
		Param86243,
		Param132049,
		Param216091,
	}

	for _, param := range params {
		t.Run(param.String(), func(t *testing.T) {
			src := New(param)
			src.Seed(4321)
		
			f, err := os.Open(filepath.Join("testdata", fmt.Sprintf("m%d.txt", param.MExp)))
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
		})
	}
}
