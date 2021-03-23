package suffixArray_test

import (
	"4420project/suffixArray"
	"4420project/suffixTree"
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	str := "alabar a la alabarda"
	str = strings.ReplaceAll(str, " ", "_")
	fmt.Println(str)
	array := suffixTree.NewSuffixTree(str).ToSuffixArray()
	sa1 := suffixArray.SuffixArray{Text: array.Text, POS: array.POS}
	sa2 := suffixArray.CreateSuffixArray(str)
	fmt.Println(sa1.POS)
	fmt.Println(sa2.POS)
}

func TestCreateSuffixArray(t *testing.T) {
	//str := "mmiissiissiippii"

}
