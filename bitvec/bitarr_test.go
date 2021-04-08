package bitvec_test

//func TestBitArr(t *testing.T) {
//	bitstr := util.GenRandomBitStr(1000000)
//	bitarr,_ := bitvec.NewBitArr(bitstr)
//	bitarr2,_ :=bitvec.NewBitArr2(bitstr)
//	for i := 0; i < len(bitstr)-32; i++ {
//		r1 := bitarr.GetValueInRange(i,i+32)
//		r2 := bitarr2.GetValueInRange(i,i+32)
//		if r1 != r2{
//			t.Error("w")
//		}
//	}
//}
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
