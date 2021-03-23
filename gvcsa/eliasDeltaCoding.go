package gvcsa

import (
	"4420project/bitvec"
	"math"
)

type EliasDeltaCoding struct {
	stream []*bitvec.BitArr
	d      *bitvec.BitArr
	e      *bitvec.BitArr
	head   []int
}

func (this *EliasDeltaCoding) Get(index int) int {
	m := this.e.Rank1(index) - 1
	s := this.e.Pred1(index)
	phi := this.head[m]
	p := this.stream[m]
	t := this.d.Pred1(index)
	arr := DecodeIntArray(p)
	startIndex := 0
	endIndex := index - s
	if t > s {
		phi = 0
		startIndex = t - s - 1
	}
	for i := startIndex; i < endIndex; i++ {
		phi += arr[i]
	}
	return phi
}

func NewEliasDeltaCoding(phi []int) *EliasDeltaCoding {
	length := len(phi)
	phiCp := make([]int, length)
	copy(phiCp, phi)
	d := bitvec.NewBitArrBySize(length)
	lastVal := phi[0]
	d.Set1(0)
	for i := 1; i < length; i++ {
		if lastVal < phi[i] {
			phi[i] = phi[i] + lastVal
			lastVal = phi[i] - lastVal
			phi[i] = phi[i] - lastVal
			phi[i] = lastVal - phi[i]
		} else {
			lastVal = phi[i]
			d.Set1(i)
		}
	}
	size := 0
	for i := 0; i < length; i++ {
		n := int(math.Log2(float64(phi[i] + 1)))
		l := int(math.Log2(float64(n + 1)))
		size += n + 2*l + 1
	}
	blockSize := 20 * int(math.Ceil(math.Log2(float64(size))))
	culSize := 0
	head := make([]int, 0)
	stream := make([]*bitvec.BitArr, 0)
	e := bitvec.NewBitArrBySize(length)
	head = append(head, phiCp[0])
	e.Set1(0)
	lastHead := 0
	for i := 1; i < length; i++ {
		n := int(math.Log2(float64(phi[i] + 1)))
		l := int(math.Log2(float64(n + 1)))
		if culSize+n+2*l+1 > blockSize {
			head = append(head, phiCp[i])
			stream = append(stream, EncodeIntArray(phi[lastHead+1:i]))
			e.Set1(i)
			culSize = 0
			lastHead = i
		} else {
			culSize += n + 2*l + 1
		}
	}
	stream = append(stream, EncodeIntArray(phi[lastHead+1:length]))
	return &EliasDeltaCoding{stream: stream, head: head, e: e, d: d}
}

func EncodeIntArray(arr []int) *bitvec.BitArr {
	length := len(arr)
	size := 0
	for i := 0; i < length; i++ {
		n := int(math.Log2(float64(arr[i] + 1)))
		l := int(math.Log2(float64(n + 1)))
		size += n + 2*l + 1
	}
	encoding := bitvec.NewBitArrBySize(size)
	currentIndex := 0
	for i := 0; i < length; i++ {
		n := int(math.Log2(float64(arr[i] + 1)))
		l := int(math.Log2(float64(n + 1)))
		encodingVal := 0
		encodingVal += (n + 1) << n
		encodingVal += (arr[i] + 1) - ((arr[i] + 1) >> n << n)
		encoding.MapValueBounded(currentIndex, currentIndex+n+2*l, uint(encodingVal))
		currentIndex = currentIndex + n + 2*l + 1
	}
	return encoding
}

func DecodeIntArray(encoding *bitvec.BitArr) []int {
	L := 0
	arr := make([]int, 0)
	for i := 0; i < encoding.Size(); {
		if encoding.Get(i) == 0 {
			L++
			i++
		} else {
			n := encoding.GetValueInRange(i, i+L) - 1
			i = i + L + 1
			val := uint(0)
			if n > 0 {
				val = encoding.GetValueInRange(i, i+int(n)-1)
			}
			val = val + 1<<n
			arr = append(arr, int(val-1))
			i = i + int(n)
			L = 0
		}
	}
	return arr
}
