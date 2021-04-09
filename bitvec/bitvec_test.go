package bitvec_test

import (
	"4420project/bitvec"
	"4420project/util"
	"math"
	"testing"
)

//func TestNewBasicBitVec(t *testing.T) {
//	//bitArr := bitvec2.NewBitArrBySize(int(math.Pow(2, 18)))
//	//bitvec := bitvec2.NewBasicBitVec(bitArr)
//	//obitvec := bitvec2.NewOneLevelBitVector(bitArr)
//	//fmt.Println(size.Of(bitvec))
//	//fmt.Println(size.Of(bitArr))
//	//fmt.Println(size.Of(obitvec))
//	bitstr := util.GenRandomBitStr(10000000)
//	arr, _ := bitvec.NewBitArr(bitstr)
//	arr2, _ := bitvec.NewBitArr2(bitstr)
//	bitarr := bitvec.NewBasicBitVec(arr)
//	bitarr2 :=bitvec.NewBasicBitVec2(arr2)
//	for i := 0; i < len(bitstr); i++ {
//		r1 := bitarr.Rank1(i)
//		r2 := bitarr2.Rank1(i)
//		if r1 != r2{
//			t.Error("w")
//		}
//	}
//}

var str = util.GenRandomBitStr(int(math.Pow(2, 20)))
var bitArr, _ = bitvec.NewBitArr(str)
var o = bitvec.NewBasicBitVec(bitArr)
var l = len(str)

func BenchmarkName(b *testing.B) {
	for i := 1; i < b.N; i++ {
		j := i % l
		r1 := o.Rank1(j)
		r2 := bitArr.Rank1(j)
		if r1 != r2 {
			print("wrong")
		}
	}
}

//func TestName(t *testing.T) {
//	str := ""
//	for i := 8; i <= 31; i++ {
//		bitArr := bitvec2.NewBitArrBySize(int(math.Pow(2, float64(i))))
//		bitvec := bitvec2.NewBasicBitVec(bitArr)
//		//obitvec := bitvec2.NewOneLevelBitVector(bitArr)
//		str += fmt.Sprint(size.Of(bitvec)) + ","
//	}
//	fmt.Println(str)
//}
