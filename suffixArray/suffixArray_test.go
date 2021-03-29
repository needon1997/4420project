package suffixArray_test

import (
	"4420project/suffixArray"
	"4420project/suffixTree"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"io/ioutil"
	"testing"
)

func TestName(t *testing.T) {
	buf, _ := ioutil.ReadFile("./test")
	str := string(buf)
	sa2 := suffixArray.CreateSuffixArray(str)
	//rlfmi := sa2.ToRLFMI()
	mtfmi := sa2.ToWTFMI()
	//sadcsa := gvcsa.MakeSADCSArray(*sa2)
	//gvcsa := gvcsa.MakeGVCSArray(*sa2)
	//fmt.Println(size.Of(str))
	//fmt.Println(size.Of(sa2))
	//fmt.Println(size.Of(rlfmi))
	fmt.Println(size.Of(mtfmi))
	//fmt.Println(size.Of(sadcsa))
	//fmt.Println(size.Of(gvcsa))
	for i := 0; i < 100000; i++ {
		_ = mtfmi.Search("AGTAGTCAGTACAGTAGTCAGTA")
	}
}

func TestCreateSuffixArray(t *testing.T) {
	//str := "mmiissiissiippii"
	buf, _ := ioutil.ReadFile("./test")
	str := string(buf)
	sa1 := suffixTree.NewSuffixTree(str).ToSuffixArray()
	sa2 := suffixArray.CreateSuffixArray(str)
	for i := 0; i < len(str); i++ {
		if sa1.POS[i] != sa2.POS[i] {
			t.Error("wrong")
		}
	}
}

func TestSuffixArray_BwtTransform(t *testing.T) {
	str := "AGTAGTCAGTAC"
	sa := suffixArray.CreateSuffixArray(str)
	fmt.Println(sa.POS)
	fmt.Println(sa.BwtTransform())
}

//var buf, _ = ioutil.ReadFile("./test")
//var str = string(buf)
//var sa2 = suffixArray.CreateSuffixArray(str)
//var rlfmi = sa2.ToWTFMI()
//var l = len(str)
//func BenchmarkName(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		rlfmi.Search("AGTAGTCAGTACAGTAGTCAGTA")
//	}
//}
