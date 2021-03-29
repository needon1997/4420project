package bitvec_test

import (
	bitvec2 "4420project/bitvec"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"io/ioutil"
	"math"
	"testing"
)

func TestNewBasicBitVec(t *testing.T) {
	bitArr := bitvec2.NewBitArrBySize(int(math.Pow(2, 18)))
	bitvec := bitvec2.NewBasicBitVec(bitArr)
	//obitvec := bitvec2.NewOneLevelBitVector(bitArr)
	fmt.Println(size.Of(bitvec))
	fmt.Println(size.Of(bitArr))
	//fmt.Println(size.Of(obitvec))
}

var buf, _ = ioutil.ReadFile("./test")
var str = string(buf)
var bitArr, _ = bitvec2.NewBitArr(str)
var o = bitvec2.NewBasicBitVec(bitArr)
var l = o.Rank1(len(str) - 1)

func BenchmarkName(b *testing.B) {
	for i := 1; i < b.N; i++ {
		o.Select1(int(math.Min(float64(i), float64(l-1))))
	}
}
