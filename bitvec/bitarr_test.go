package bitvec_test

import (
	"4420project/bitvec"
	"fmt"
	"github.com/DmitriyVTitov/size"
	"testing"
)

func TestBitArr(t *testing.T) {
	bitArr := bitvec.NewBitArrBySize(110016)
	fmt.Println(size.Of(bitvec.NewBasicBitVec(bitArr)))
}
