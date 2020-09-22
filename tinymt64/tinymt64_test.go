package tinymt64_test

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/shogo82148/randsrc/tinymt64"
)

func BenchmarkInt63(b *testing.B) {
	src := tinymt64.New(0, 0, 0, [2]uint64{})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Int63())
	}
}

func BenchmarkUint64(b *testing.B) {
	src := tinymt64.New(0, 0, 0, [2]uint64{})
	for i := 0; i < b.N; i++ {
		runtime.KeepAlive(src.Uint64())
	}
}

func ExampleSource_Uint64() {
	src := tinymt64.New(0xfa051f40, 0xffd0fff4, 0x58d02ffeffbfffbc, [2]uint64{})
	src.Seed(1)
	for i := 0; i < 30; i++ {
		fmt.Println(src.Uint64())
	}
	//Output:
	// 15503804787016557143
	// 17280942441431881838
	// 2177846447079362065
	// 10087979609567186558
	// 8925138365609588954
	// 13030236470185662861
	// 4821755207395923002
	// 11414418928600017220
	// 18168456707151075513
	// 1749899882787913913
	// 2383809859898491614
	// 4819668342796295952
	// 11996915412652201592
	// 11312565842793520524
	// 995000466268691999
	// 6363016470553061398
	// 7460106683467501926
	// 981478760989475592
	// 11852898451934348777
	// 5976355772385089998
	// 16662491692959689977
	// 4997134580858653476
	// 11142084553658001518
	// 12405136656253403414
	// 10700258834832712655
	// 13440132573874649640
	// 15190104899818839732
	// 14179849157427519166
	// 10328306841423370385
	// 9266343271776906817
}

func ExampleSource_SeedBySlice() {
	src := tinymt64.New(0xfa051f40, 0xffd0fff4, 0x58d02ffeffbfffbc, [2]uint64{})
	src.SeedBySlice([]uint64{1})
	for i := 0; i < 30; i++ {
		fmt.Println(src.Uint64())
	}
	//Output:
	// 2316304586286922237
	// 15094277089150361724
	// 5685675787316092711
	// 15229481068059623199
	// 4714098425347676722
	// 16281862982583854132
	// 3901922025624662484
	// 5886484389080126014
	// 16107583395258923453
	// 13952088220369493459
	// 17758435316338264754
	// 2351799565271811353
	// 12362529980853249542
	// 1719516909033106250
	// 8766952554732792269
	// 7859523628104690493
	// 15389348425598624967
	// 5147268256773563271
	// 9499111560078684970
	// 667293060984396585
	// 16412518715911243540
	// 4644561915126619944
	// 7147182560776836637
	// 1588726635616164641
	// 14118193191231902733
	// 10534117574818039474
	// 5944505171977344673
	// 443288919934395040
	// 1633068730058384525
	// 17771926205819909233
}
