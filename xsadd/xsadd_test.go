package xsadd

import (
	"runtime"
	"testing"
)

func BenchmarkInt63(b *testing.B) {
	src := New([4]uint32{})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := New([4]uint32{})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func BenchmarkUint32(b *testing.B) {
	src := New([4]uint32{})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint32())
	}
}

func TestSeed(t *testing.T) {
	var src Source
	src.Seed(1234)
	state := [4]uint32{0xbfb2c4f3, 0xe5f5b22e, 0x8b9a4397, 0x0156d240}
	if src.state != state {
		t.Errorf("invalid state: want %#v, got %#v", state, src.state)
	}

	wants := [...]uint32{
		1823491521, 1658333335, 1467485721, 45623648,
		3336175492, 2561136018, 181953608, 768231638,
		3747468990, 633754442, 1317015417, 2329323117,
		688642499, 1053686614, 1029905208, 3711673957,
		2701869769, 695757698, 3819984643, 1221024953,
		110368470, 2794248395, 2962485574, 3345205107,
		592707216, 1730979969, 2620763022, 670475981,
		1891156367, 3882783688, 1913420887, 1592951790,
		2760991171, 1168232321, 1650237229, 2083267498,
		2743918768, 3876980974, 2059187728, 3236392632,
	}
	for i, want := range wants {
		got := src.Uint32()
		if got != want {
			t.Errorf("mismatch %d: want %d, got %d", i, want, got)
		}
	}
}

func TestSeedBySlice(t *testing.T) {
	var src Source
	src.SeedBySlice([]uint32{0x0a, 0x0b, 0x0c, 0x0d})
	state := [4]uint32{0x76648e9b, 0x6fecb209, 0x3ac0fe4c, 0x54f1f628}
	if src.state != state {
		t.Errorf("invalid state: want %#v, got %#v", state, src.state)
	}

	wants := [...]uint32{
		0x138a38f9, 0xb396fa84, 0xa55a2ee8, 0x24b7ed06,
		0xf0bae2fe, 0xd8ace1a7, 0xd4b09a3f, 0xd7fcf441,
		0xfc55ee1b, 0x5b4ab585, 0xd4bf254b, 0x5b0f77ba,
		0x31161b97, 0xb21ccc3b, 0xab418bfb, 0x4cc8476a,
		0x06a1a28f, 0xcb1f50c6, 0xf0ba2ed3, 0x7907f372,
		0x3256d76c, 0xd843e864, 0xd63a60b7, 0xeff88358,
		0xddc3b083, 0xb5734b65, 0xf08d644d, 0xe5f6c809,
		0x95bf2ae3, 0xe5867758, 0xf260d462, 0x39d244dc,
		0xb9fbb8d7, 0x63e8f3d9, 0xb34ea936, 0x8fe4ee75,
		0x8803c8f1, 0xd74e420e, 0xa5c14d22, 0x20be253f,
	}
	for i, want := range wants {
		got := src.Uint32()
		if got != want {
			t.Errorf("mismatch %d: want %d, got %d", i, want, got)
		}
	}
}
