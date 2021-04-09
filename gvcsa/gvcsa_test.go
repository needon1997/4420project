package gvcsa_test

import (
	"4420project/gvcsa"
	"4420project/suffixArray"
	"4420project/suffixTree"
	"4420project/util"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"io/ioutil"
	"testing"
)

func TestMakeGVCSArray(t *testing.T) {
	buf, _ := ioutil.ReadFile("./test")
	str := string(buf)
	tree := suffixTree.NewSuffixTree(str)
	array := tree.ToSuffixArray()
	sa := suffixArray.SuffixArray{Text: array.Text, POS: array.POS}
	fmi := sa.ToWTFMI()
	rlfmi := sa.ToRLFMI()
	gvcsArray := gvcsa.MakeGVCSArray(sa)
	sadcsArray := gvcsa.MakeSADCSArray(sa)
	fmt.Println(size.Of(str))
	fmt.Println(size.Of(tree))
	fmt.Println(size.Of(sa))
	fmt.Println(size.Of(gvcsArray))
	fmt.Println(size.Of(sadcsArray.Phi[0]))
	fmt.Println(size.Of(fmi))
	fmt.Println(size.Of(rlfmi))
	//for i := 0; i < 500; i++ {
	//	r1 := gvcsArray.Search(str[0+i : 50+i])
	//	r2 := sadcsArray.Search(str[0+i : 50+i])
	//	r3 := rlfmi.Search(str[0+i : 50+i])
	//	r4 := fmi.Search(str[0+i : 50+i])
	//	if true{
	//		fmt.Println(r1,r2,r3,r4)
	//	}
	//}
	//fmt.Println(time1)
	//fmt.Println(time2)
}

func TestEncodeIntArray(t *testing.T) {
	str := util.GenRandomStr(15, 3)
	fmt.Println(size.Of(str))
	sa := suffixArray.CreateSuffixArray(str)
	sad := gvcsa.MakeSADCSArray(*sa)
	for i := 0; i < 10; i++ {
		r1 := sad.Substring(i, i+3)
		r2 := sa.Text[i : i+3]
		if r1 != r2 {
			t.Error("wrong")
		}
	}
}
