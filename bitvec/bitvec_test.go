package bitvec_test

import (
	bitvec2 "4420project/bitvec"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"math"
	"testing"
)

func TestNewBasicBitVec(t *testing.T) {
	bitArr := bitvec2.NewBitArrBySize(int(math.Pow(2, 20)))
	bitvec := bitvec2.NewBasicBitVec(bitArr)
	fmt.Println(size.Of(bitvec))
	fmt.Println(size.Of(bitArr))
}

func TestBasicBitVec(t *testing.T) {
	str := "00000000000000100000000101101101000001100100010010111011100011011111000001001101111000100100100010000111000110100110100100111010100101110010011101010110011101101110101111000100000100010101100011100110110101101001100001111000010111101010101110001100010100000101111000111000101011000001101110101001011000100010101111101101110111000100000001111001110011001011101101100011000100001101100100111000101001001011100110101110101111011100101111011011010000000010011000110001011001101001111100110100100000010000"
	bitvec, _ := bitvec2.NewBasicBitVecFromString(str)
	bitArr, _ := bitvec2.NewBitArr(str)
	for i := 0; i < len(str); i++ {
		r1 := bitvec.Rank1(i)
		r2 := bitArr.Rank1(i)
		if r1 != r2 {
			t.Error("wrong")
		}
		i1 := bitvec.Select1(r1)
		i2 := bitArr.Select1(r2)
		if i1 != i2 {
			t.Error("wrong")
		}
	}
}
