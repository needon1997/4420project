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
	rlfmi := sa2.ToRLFMI()
	mtfmi := sa2.ToWTFMI()
	fmt.Println(size.Of(str))
	fmt.Println(size.Of(rlfmi))
	fmt.Println(size.Of(mtfmi))
	for i := 0; i < 300; i++ {
		R1 := rlfmi.Search(str[0+i : 500+i])
		R2 := mtfmi.Search(str[0+i : 500+i])
		if R1 != R2 {
			t.Error("wrong")
		}
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
