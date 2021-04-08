package bitvec

import (
	"fmt"
	"math"
)

const (
	subBlockSize = 32
)

var popC = GenPopCount()

func GenPopCount() []byte {
	table := make([]byte, 256)
	for i := 0; i <= 255; i++ {
		bitArr := NewBitArrBySize(8)
		bitArr.MapValueBounded(0, 7, uint(i))
		table[i] = byte(bitArr.Rank1(7))
	}
	return table
}

type BasicBitVector struct {
	length              int
	blockSize           int
	blockNum            int
	blockRankBitsNum    int
	subBlockSize        int
	subBlockNum         int
	subBlockRankBitsNum int
	popc                []byte
	blockRank           *BitArr
	subBlockRank        *BitArr
	bitArr              *BitArr
}

func (this *BasicBitVector) String() string {
	str := ""
	for i := 0; i < this.length; i++ {
		str += fmt.Sprint(this.Get(i))
	}
	return str
}
func (this *BasicBitVector) Get(i int) uint8 {
	return this.bitArr.Get(i)
}
func (this *BasicBitVector) Size() int {
	return this.length
}
func (this *BasicBitVector) GetValueInRange(start int, end int) uint {
	return this.bitArr.GetValueInRange(start, end)
}
func (this *BasicBitVector) Pred1(i int) int {
	return this.Select1(this.Rank1(i))
}
func (this *BasicBitVector) Select1(j int) int {
	s := 0
	e := this.blockNum - 1
	if j > this.Rank1(this.length-1) {
		panic("invalid j")
	}
	if j == 0 {
		return -1
	}
	i := (s + e) / 2
	l := uint(0)
	for i >= 1 {
		l = this.blockRank.GetValueInRange((i-1)*this.blockRankBitsNum, i*this.blockRankBitsNum-1)
		r := this.blockRank.GetValueInRange(i*this.blockRankBitsNum, (i+1)*this.blockRankBitsNum-1)
		if j > int(l) && j <= int(r) {
			break
		} else if j <= int(l) {
			e = i - 1
		} else if j > int(r) {
			s = i + 1
		}
		i = (s + e) / 2
	}
	if i == 0 {
		l = 0
	}
	j = j - int(l)
	s = 0
	for {
		b := this.bitArr.arr[i*this.subBlockNum*4+s]
		r := int(this.popc[b])
		if j-r <= 0 {
			break
		} else {
			j = j - r
		}
		s++
	}
	for k := 0; k < 8; k++ {
		bit := this.bitArr.Get(8*(i*this.subBlockNum*4+s) + k)
		if bit == 1 {
			j = j - 1
		}
		if j == 0 {
			return 8*(i*this.subBlockNum*4+s) + k
		}
	}
	return j
}

func (this *BasicBitVector) Rank0(index int) int {
	return index + 1 - this.Rank1(index)
}
func (this *BasicBitVector) Rank1(index int) int {
	b := (index + 1) / this.blockSize
	j := (index + 1) % this.blockSize
	c := j / this.subBlockSize
	k := j % this.subBlockSize
	by := k / 8
	r := k % 8
	var rank1 uint
	if b == 0 {
		rank1 = 0
	} else {
		rank1 = this.blockRank.GetValueInRange((b-1)*this.blockRankBitsNum, b*this.blockRankBitsNum-1)
	}
	var rank2 uint
	if c == 0 {
		rank2 = 0
	} else {
		rank2 = this.subBlockRank.GetValueInRange((b*this.subBlockNum+c-1)*this.subBlockRankBitsNum, (b*this.subBlockNum+c)*this.subBlockRankBitsNum-1)
	}
	var rank3 uint = 0
	for i := 0; i < by; i++ {
		rank3 += uint(this.popc[uint8(this.bitArr.arr[c+b*this.subBlockNum]>>(24-i*8))])
	}
	rank3 += uint(this.popc[uint8(this.bitArr.arr[c+b*this.subBlockNum]>>(24-by*8))>>(8-r)])
	return int(rank1) + int(rank2) + int(rank3)
}
func NewBasicBitVec(b *BitArr) *BasicBitVector {
	blockSize := subBlockSize * int(math.Log2(float64(b.length)))
	blockNum := int(math.Ceil(float64(b.length) / float64(blockSize)))
	blockRankBitsNum := int(math.Ceil(math.Log2(float64(b.length + 1))))
	blockRank := NewBitArrBySize(blockNum * blockRankBitsNum)
	subBlockNum := int(math.Log2(float64(b.length)))
	subBlockRankBitsNum := int(math.Ceil(math.Log2(float64(blockSize + 1))))
	subBlockRank := NewBitArrBySize(int(math.Ceil(float64(b.length)/float64(subBlockSize))) * subBlockRankBitsNum)
	subBlockIndex := 0
	prevBlockRank := 0
	for i := 1; i <= blockNum; i++ {
		var blockRankVal int
		if i*blockSize < b.length {
			blockRankVal = b.RangeRank1((i-1)*blockSize, i*blockSize-1) + prevBlockRank
		} else {
			blockRankVal = b.RangeRank1((i-1)*blockSize, b.length-1) + prevBlockRank
		}
		blockRank.MapValueBounded((i-1)*blockRankBitsNum, i*blockRankBitsNum-1, uint(blockRankVal))
		prevSubBlockRank := 0
		subBlockLoopEnded := false
		for j := 1; j <= subBlockNum && !subBlockLoopEnded; j++ {
			var subBlockRankVal int
			if j*subBlockSize+(i-1)*blockSize < b.length {
				subBlockRankVal = b.RangeRank1((i-1)*blockSize+(j-1)*subBlockSize, (i-1)*blockSize+j*subBlockSize-1) + prevSubBlockRank
			} else {
				subBlockRankVal = b.RangeRank1((i-1)*blockSize+(j-1)*subBlockSize, b.length-1) + prevSubBlockRank
				subBlockLoopEnded = true
			}
			subBlockRank.MapValueBounded(subBlockIndex*subBlockRankBitsNum, (1+subBlockIndex)*subBlockRankBitsNum-1, uint(subBlockRankVal))
			subBlockIndex += 1
			prevSubBlockRank = subBlockRankVal
		}
		prevBlockRank = blockRankVal
	}
	bitv := &BasicBitVector{}
	bitv.blockSize = blockSize
	bitv.blockNum = blockNum
	bitv.blockRankBitsNum = blockRankBitsNum
	bitv.blockRank = blockRank
	bitv.subBlockSize = subBlockSize
	bitv.subBlockNum = subBlockNum
	bitv.subBlockRankBitsNum = subBlockRankBitsNum
	bitv.subBlockRank = subBlockRank
	bitv.bitArr = b
	bitv.length = b.length
	bitv.popc = popC
	return bitv
}

func GenPopCount2() []byte {
	table := make([]byte, 256)
	for i := 0; i <= 255; i++ {
		bitArr := NewBitArrBySize(8)
		bitArr.MapValueBounded(0, 7, uint(i))
		table[i] = byte(bitArr.Rank1(7))
	}
	return table
}
