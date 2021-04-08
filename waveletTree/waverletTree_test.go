package waveletTree_test

import (
	"4420project/util"
	"4420project/waveletTree"
	"testing"
)

func TestNewWaveletTree(t *testing.T) {
	var str = util.GenRandomStr(10000000, 4)
	var tree1 = waveletTree.NewWaveletTree(str, []byte("abcd"))
	var tree2 = waveletTree.NewWaveletTree2(str, []byte("abcd"))
	for i := 0; i < 100; i++ {
		r1 := tree1.Rank([]byte("a")[0], 10000)
		r2 := tree2.Rank([]byte("a")[0], 10000)
		if r1 != r2 {
			t.Error("w")
		}
	}
}

var str = util.GenRandomStr(10000000, 4)
var tree1 = waveletTree.NewWaveletTree(str, []byte("abcd"))

//var tree2 = waveletTree.NewWaveletTree2(str, []byte("abcd"))
func BenchmarkName(b *testing.B) {
	for i := 1; i < b.N; i++ {
		j := i % 10000000
		tree1.Rank([]byte("d")[0], j)
	}
}
