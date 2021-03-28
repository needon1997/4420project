package waveletTree_test

import (
	"4420project/waveletTree"
	"io/ioutil"
	"testing"
)

func TestNewWaveletTree(t *testing.T) {
	buf, _ := ioutil.ReadFile("./test")
	str := string(buf)

	tree := waveletTree.NewWaveletTree(str, []string{"A", "C", "T", "G"})
	for i := 0; i < len(str); i++ {
		if tree.Get(i) != str[i:i+1] {
			t.Error("w")
		}
	}
}
