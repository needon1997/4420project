package bitvec

//
//
//
//import (
//	"math"
//	"unsafe"
//)
//
//type BitVector interface {
//	Rank0(index int) int
//	Rank1(index int) int
//	Select0(index int) int
//	Select1(index int) int
//}
//
//type BasicBitVector struct {
//	length                      int
//	blockSize                   int
//	subBlockSize                int
//	blockRankBitsNum            int
//	subBlockRankBitsNum         int
//	subBlockBitRankBitsNum      int
//	subBlockBitRankIndexBitsNum int
//	subBlockNum                 int
//	superBlockRank              *BitArr
//	subBlockRank                *BitArr
//	subBlockBitRank             *BitArr
//	subBlockBitRankIndex        *BitArr
//	bitArr                      *BitArr
//}
//
//func (this *BasicBitVector) Get(i int) uint8 {
//	return this.bitArr.Get(i)
//}
//func (this *BasicBitVector) Size() int {
//	return this.length
//}
//func (this *BasicBitVector) GetValueInRange(start int, end int) uint {
//	return this.bitArr.GetValueInRange(start, end)
//}
//func (this *BasicBitVector) SizeOf() uintptr {
//	return unsafe.Sizeof(*this) + unsafe.Sizeof(*this.superBlockRank) + unsafe.Sizeof(*this.subBlockBitRank) + unsafe.Sizeof(*this.subBlockRank) + unsafe.Sizeof(*this.subBlockBitRankIndex)
//}
//func (this *BasicBitVector) Pred1(i int) int {
//	return this.Select1(this.Rank1(i))
//}
//func (this *BasicBitVector) Select1(j int) int {
//	s := 0
//	e := this.length
//	if j > this.Rank1(this.length-1) {
//		panic("invalid j")
//	}
//	if j == 0 {
//		return -1
//	}
//	i := (s + e) / 2
//	for true {
//		k := this.Rank1(i)
//		if k > j {
//			e = i
//		} else if k < j {
//			s = i
//		} else {
//			if i == 0 || this.Rank1(i-1) == j-1 {
//				return i
//			} else {
//				e = i
//			}
//		}
//		i = (s + e) / 2
//	}
//	return -1
//}
//
//func (this *BasicBitVector) Rank0(index int) int {
//	return index + 1 - this.Rank1(index)
//}
//func (this *BasicBitVector) Rank1(index int) int {
//	b := index / this.blockSize
//	j := index % this.blockSize
//	c := j / this.subBlockSize
//	k := j % this.subBlockSize
//	var rank1, rank2 uint
//	if b == 0 {
//		rank1 = 0
//	} else {
//		rank1 = this.superBlockRank.GetValueInRange((b-1)*this.blockRankBitsNum, b*this.blockRankBitsNum-1)
//	}
//	if c == 0 {
//		rank2 = 0
//	} else {
//		rank2 = this.subBlockRank.GetValueInRange((b*this.subBlockNum+c-1)*this.subBlockRankBitsNum, (b*this.subBlockNum+c)*this.subBlockRankBitsNum-1)
//	}
//	i := b*this.blockSize + c*this.subBlockSize
//	l := int(math.Min(float64(b*this.blockSize+(c+1)*this.subBlockSize), math.Min(float64((b+1)*this.blockSize), float64(this.length))))
//	o := this.subBlockSize - (l - i)
//	val := this.bitArr.GetValueInRange(i, l-1) << o
//	rank3 := this.subBlockBitRank.GetValueInRange((int(val)*this.subBlockSize+k)*this.subBlockBitRankBitsNum, (int(val)*this.subBlockSize+k+1)*this.subBlockBitRankBitsNum-1)
//	return int(rank1 + rank2 + rank3)
//}
//
//func NewBasicBitVec(bitArr *BitArr) *BasicBitVector {
//	arrSize := bitArr.length
//	blockRankBitsNum := int(math.Ceil(math.Log2(float64(arrSize + 1))))
//	blockSize := int(math.Ceil(math.Log2(float64(arrSize)) * math.Log2(float64(arrSize))))
//	blockNum := int(math.Ceil(float64(arrSize) / float64(blockSize)))
//
//	subBlockRankBitsNum := int(math.Ceil(math.Log2(float64(blockSize + 1))))
//	subBlockSize := int(math.Ceil(0.5 * math.Log2(float64(arrSize))))
//	subBlockNum := int(math.Ceil(float64(blockSize) / float64(subBlockSize)))
//	subBlockBitRankBitsNum := int(math.Ceil(math.Log2(float64(subBlockSize + 1))))
//
//	sqrtN := int(math.Ceil(math.Pow(2, float64(subBlockSize))))
//	bv := &BasicBitVector{length: arrSize}
//	bv.blockSize = blockSize
//	bv.subBlockSize = subBlockSize
//	bv.superBlockRank = NewBitArrBySize(blockRankBitsNum * blockNum)
//	bv.subBlockRank = NewBitArrBySize(blockNum * subBlockNum * subBlockRankBitsNum)
//	bv.subBlockBitRank = NewBitArrBySize(sqrtN * subBlockSize * subBlockBitRankBitsNum)
//	bv.blockRankBitsNum = blockRankBitsNum
//	bv.subBlockRankBitsNum = subBlockRankBitsNum
//	bv.subBlockBitRankBitsNum = subBlockBitRankBitsNum
//	bv.subBlockNum = subBlockNum
//	bv.bitArr = bitArr
//	subBlockIndex := 0
//	prevBlockRank := 0
//	//initialize the sqrt(n)  unique bit string with length log n/2
//	for i := 0; i < sqrtN; i++ {
//		temp := NewBitArrBySize(subBlockSize)
//		temp.MapValueBounded(0, subBlockSize-1, uint(i))
//		for j := 0; j < subBlockSize; j++ {
//			rank := temp.Rank1(j)
//			bv.subBlockBitRank.MapValueBounded((i*subBlockSize+j)*subBlockBitRankBitsNum, (i*subBlockSize+j+1)*subBlockBitRankBitsNum-1, uint(rank))
//		}
//	}
//	for i := 1; i <= blockNum; i++ {
//		var blockRankVal int
//		if i*blockSize < arrSize {
//			blockRankVal = bitArr.RangeRank1((i-1)*blockSize, i*blockSize-1) + prevBlockRank
//		} else {
//			blockRankVal = bitArr.RangeRank1((i-1)*blockSize, arrSize-1) + prevBlockRank
//		}
//		bv.superBlockRank.MapValueBounded((i-1)*blockRankBitsNum, i*blockRankBitsNum-1, uint(blockRankVal))
//		prevSubBlockRank := 0
//		subBlockLoopEnded := false
//		for j := 1; j <= subBlockNum && !subBlockLoopEnded; j++ {
//			var subBlockRankVal int
//			if j*subBlockSize+(i-1)*blockSize < arrSize {
//				if j*subBlockSize < blockSize {
//					subBlockRankVal = bitArr.RangeRank1((i-1)*blockSize+(j-1)*subBlockSize, (i-1)*blockSize+j*subBlockSize-1) + prevSubBlockRank
//				} else {
//					subBlockRankVal = bitArr.RangeRank1((i-1)*blockSize+(j-1)*subBlockSize, i*blockSize-1) + prevSubBlockRank
//				}
//			} else {
//				subBlockRankVal = bitArr.RangeRank1((i-1)*blockSize+(j-1)*subBlockSize, arrSize-1)
//				subBlockLoopEnded = true
//			}
//			bv.subBlockRank.MapValueBounded(subBlockIndex*subBlockRankBitsNum, (1+subBlockIndex)*subBlockRankBitsNum-1, uint(subBlockRankVal))
//			subBlockIndex += 1
//			prevSubBlockRank = subBlockRankVal
//		}
//		prevBlockRank = blockRankVal
//	}
//	return bv
//}
//
//func NewBasicBitVecFromString(bitstring string) (*BasicBitVector, error) {
//	bitArr, err := NewBitArr(bitstring)
//	if err != nil {
//		return nil, err
//	}
//	return NewBasicBitVec(bitArr), nil
//}
