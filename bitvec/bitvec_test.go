package bitvec_test

import (
	bitvec2 "4420project/bitvec"
	"fmt"
	"github.com/DmitriyVTitov/size"
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

//
func TestBasicBitVec(t *testing.T) {
	str := "1110111110000111111111"
	bitArr, _ := bitvec2.NewBitArr(str)
	o := bitvec2.NewBasicBitVec(bitArr)
	for i := 0; i < len(str); i++ {
		R1 := bitArr.Rank0(i)
		R2 := o.Rank0(i)
		if R1 != R2 {
			t.Error("w")
		}
	}
}
