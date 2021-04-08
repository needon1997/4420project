package waveletTree_test

import (
	"4420project/util"
	"4420project/waveletTree"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"testing"
)

func TestNewWaveletTree(t *testing.T) {
	var str = util.GenRandomStr(100000, 4)
	var tree1 = waveletTree.NewWaveletTree(str, []byte("abcd"))
	var tree2 = waveletTree.NewWaveletTree2(str, []byte("abcd"))
	fmt.Println(size.Of(tree1))
	fmt.Println(size.Of(tree2))
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
