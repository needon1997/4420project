package bitvec_test

import (
	"4420project/bitvec"
	"testing"
)

func TestBitArr(t *testing.T) {
	//bitstr := util.GenRandomBitStr(1000000)
	//bitarr,_ := bitvec.NewBitArr(bitstr)
	//bitarr2,_ :=bitvec.NewBitArr2(bitstr)
	for i := 0; i < 1000000; i++ {
		arr := bitvec.ToBitArr(uint(i))
		s1 := arr.String()
		arr2 := bitvec.ToBitArr2(uint(i))
		s2 := arr2.String()
		if s1 != s2 {
			t.Error("w")
		}
	}
}

//var bitstr = util.GenRandomBitStr(1000000)
//var bitarr,_ = bitvec.NewBitArr(bitstr)
//var bitarr2,_ =bitvec.NewBitArr2(bitstr)
//var l1 = len(bitstr)
//func BenchmarkName111(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		j := i % (l1-32)
//		bitarr2.GetValueInRange(j,j+32)
//	}
//}
