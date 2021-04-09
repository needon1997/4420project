package compactArr

import "4420project/bitvec"

type CompactArr struct {
	elementsSize int
	size         int
	arr          bitvec.BitArr
}

func (this *CompactArr) Set(index int, value uint) {
	if index > this.size-1 || index < 0 {
		panic("Array Index out of Bound")
	}
	this.arr.MapValueBounded(index*this.elementsSize, (index+1)*this.elementsSize-1, value)
}

func (this *CompactArr) Get(index int) uint {
	if index > this.size-1 || index < 0 {
		panic("Array Index out of Bound")
	}
	return this.arr.GetValueInRange(index*this.elementsSize, (index+1)*this.elementsSize-1)
}
func NewCompactArr(bits int, size int) *CompactArr {
	arr := bitvec.NewBitArrBySize(bits * size)
	return &CompactArr{arr: *arr, size: size, elementsSize: bits}
}
