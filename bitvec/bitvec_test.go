package bitvec_test

import (
	"4420project/bitvec"
	"4420project/util"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"math"
	"testing"
)

func TestNewBasicBitVec(t *testing.T) {
	bitstr := util.GenRandomBitStr(int(math.Pow(2, 5)))
	bitarr, _ := bitvec.NewBitArr(bitstr)
	bitvec := bitvec.NewBasicBitVec(bitarr)
	fmt.Println(size.Of(bitvec))
	fmt.Println(size.Of(bitarr))
	for i := 0; i < len(bitstr); i++ {
		r1 := bitarr.Rank1(i)
		r2 := bitvec.Rank1(i)
		if r1 != r2 {
			t.Error("w")
		}
	}
}

//var str = util.GenRandomBitStr(int(math.Pow(2, 20)))
//var bitArr, _ = bitvec.NewBitArr(str)
//var o = bitvec.NewBasicBitVec(bitArr)
//var l = o.Rank1(len(str) - 1)
//
//func BenchmarkName(b *testing.B) {
//	for i := 1; i < b.N; i++ {
//		j := i % l
//		_ = o.Rank1(j)
//	}
//}
