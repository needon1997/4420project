package bitvec

import (
	"errors"
	"fmt"
	"math"
)

type BitArr struct {
	arr    []uint32
	length int
}

func NewBitArrBySize(n int) *BitArr {
	size := int(math.Ceil(float64(n) / float64(32)))
	barr := BitArr{arr: make([]uint32, size), length: n}
	return &barr
}

func NewBitArr(bitString string) (*BitArr, error) {
	l := len(bitString)
	size := int(math.Ceil(float64(l) / float64(32)))
	barr := BitArr{arr: make([]uint32, size), length: l}
	for i := 0; i < l; i++ {
		if bitString[i] == 48 {
			barr.Set0(i)
		} else if bitString[i] == 49 {
			barr.Set1(i)
		} else {
			return nil, errors.New("wrong bit string format")
		}
	}
	return &barr, nil
}

func Value(barr BitArr) uint {
	var val uint = 1
	l := len(barr.arr)
	var sum uint = 0
	for i := l - 1; i >= 0; i-- {
		sum += uint(barr.arr[i]) * val
		val *= val << 8
	}
	return sum
}
func ToBitArr(val uint) BitArr {
	if val == 0 {
		return BitArr{arr: make([]uint32, 1), length: 8}
	}
	if val < 0 {
		panic("not support")
	}
	cval := val
	blockSize := 0
	for val != 0 {
		val = val >> 32
		blockSize += 1
	}
	arr := make([]uint32, blockSize)
	for i := blockSize - 1; i >= 0; i-- {
		arr[i] = uint32(cval - cval>>32<<32)
		cval = cval >> 32
	}
	return BitArr{arr: arr, length: blockSize * 32}
}
func (this *BitArr) String() string {
	str := ""
	for i := 0; i < this.length; i++ {
		str += fmt.Sprint(this.Get(i))
	}
	return str
}
func (this *BitArr) GetValueInRange(start, end int) uint {
	if end < start {
		panic("end < start")
	} else if start < 0 {
		panic("start < 0")
	} else if end >= this.length {
		panic("index out of bound")
	}
	startSuperIndex := start / 32
	startIndex := start % 32
	endSuperIndex := end / 32
	endIndex := end % 32
	var val uint = 0
	if startSuperIndex != endSuperIndex {
		val += uint(this.arr[endSuperIndex]) >> (32 - (endIndex + 1))
		for i := 0; i < endSuperIndex-startSuperIndex-1; i++ {
			val += uint(this.arr[endSuperIndex-i-1]) << (endIndex + 1 + i*32)
		}
		val += uint(this.arr[startSuperIndex]<<startIndex>>startIndex) << (endIndex + 1 + (endSuperIndex-startSuperIndex-1)*32)
	} else {
		val += uint(this.arr[startSuperIndex] << startIndex >> (startIndex + (32 - (endIndex + 1))))
	}
	return val
}
func (this *BitArr) MapValueBounded(start, end int, val uint) {
	if end < start {
		panic("end < start")
	}
	ba := ToBitArr(val)
	for i := 0; i <= end-start; i++ {
		if ba.length-1-i < 0 {
			this.Set0(end - i)
			continue
		}
		v := ba.Get(ba.length - 1 - i)
		if v == 1 {
			this.Set1(end - i)
		} else {
			this.Set0(end - i)
		}
	}
}
func (this *BitArr) Get(i int) uint8 {
	if i >= this.length || i < 0 {
		fmt.Println("error")
		panic("index out of bound")
	}
	superIndex := i / 32
	index := i % 32
	block := this.arr[superIndex]
	ops := uint32(1) << (31 - index)
	result := block & ops >> (31 - index)
	return uint8(result)

}
func (this *BitArr) Set1(i int) {
	if i >= this.length || i < 0 {
		panic("index out of bound")
	}
	superIndex := i / 32
	index := i % 32
	ops := uint32(1) << (31 - index)
	this.arr[superIndex] = this.arr[superIndex] | ops
}
func (this *BitArr) Set0(i int) {
	if i >= this.length || i < 0 {
		panic("index out of bound")
	}
	superIndex := i / 32
	index := i % 32
	ops := uint32(uint32(1<<32-1) - uint32(1)<<(31-index))
	this.arr[superIndex] = this.arr[superIndex] & ops
}

func (this *BitArr) Rank0(index int) int {
	rank := 0
	for i := 0; i <= index; i++ {
		if this.Get(i) == 0 {
			rank += 1
		}
	}
	return rank
}
func (this *BitArr) RangeRank0(start int, end int) int {
	rank := 0
	for i := start; i <= end; i++ {
		if this.Get(i) == 0 {
			rank += 1
		}
	}
	return rank
}
func (this *BitArr) RangeRank1(start int, end int) int {
	return end - start + 1 - this.RangeRank0(start, end)
}
func (this *BitArr) Rank1(index int) int {
	return index + 1 - this.Rank0(index)
}
func (this *BitArr) Select1(j int) int {
	cul := 0
	if j == 0 {
		return 0
	}
	for i := 0; i < this.length; i++ {
		if this.Get(i) == 1 {
			cul++
			if cul == j {
				return i
			}
		}
	}
	return -1
}
func (this *BitArr) Pred1(index int) int {
	for i := index; i >= 0; i-- {
		if this.Get(i) == 1 {
			return i
		}
	}
	return -1
}
func (this *BitArr) Size() int {
	return this.length
}
