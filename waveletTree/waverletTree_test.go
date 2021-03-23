package waveletTree_test

import (
	"4420project/waveletTree"
	"testing"
)

func TestNewWaveletTree(t *testing.T) {
	tree := waveletTree.NewWaveletTree("abcabxabcd", []string{"a", "b", "c", "d", "x"})
	println(tree.Rank("a", 8))
}
