package suffixArray

import (
	"4420project/bitvec"
)

type subString struct {
	name  uint32
	start uint32
	end   uint32
}

func CreateSuffixArray(text string) *SuffixArray {
	text = text + string(byte(1))
	n := len(text)
	modText := make([]uint32, len(text))
	for i := 0; i < n; i++ {
		modText[i] = uint32(text[i])
	}
	sa := createSuffixArrayRecur(modText)
	return &SuffixArray{Text: text, POS: sa}
}
func createSuffixArrayRecur(text []uint32) []int {
	bucket := getBucket(text)
	t := getType(text)
	p := findLMSPtrArray(t)
	sa := putBucket(text, p, bucket, t)
	nameP(p, sa, text, t)
	unique, sa1 := checkUnique(p)
	if !unique {
		text1 := make([]uint32, 0)
		n := len(p)
		for i := 0; i < n; i++ {
			if p[i] != nil {
				text1 = append(text1, p[i].name)
			}
		}
		newP := make([]*subString, len(text1))
		j := 0
		for i := 0; i < n; i++ {
			if p[i] != nil {
				newP[j] = p[i]
				j++
			}
		}
		p = newP
		sa1 = createSuffixArrayRecur(text1)
	}
	sa1Length := len(sa1)
	p1 := make([]*subString, sa1Length)
	for i := 0; i < sa1Length; i++ {
		p1[i] = p[sa1[i]]
	}
	sa = putBucket(text, p1, bucket, t)
	return sa
}
func getBucket(text []uint32) []uint32 {
	n := len(text)
	var max uint32 = 0
	for i := 0; i < n; i++ {
		if text[i] > max {
			max = text[i]
		}
	}
	bucket := make([]uint32, max+1)
	for i := 0; i < n; i++ {
		bucket[text[i]] = bucket[text[i]] + 1
	}
	var culsum uint32 = 0
	var i uint32 = 0
	for ; i < max+1; i++ {
		culsum = bucket[i] + culsum
		bucket[i] = culsum - 1
	}

	return bucket
}
func getType(text []uint32) *bitvec.BitArr {
	t := bitvec.NewBitArrBySize(len(text))
	stype := true
	t.Set1(len(text) - 1)
	for i := len(text) - 2; i >= 0; i-- {
		if text[i] != text[i+1] {
			if text[i] < text[i+1] {
				stype = true
				t.Set1(i)
			} else {
				stype = false
			}
		} else {
			if stype {
				t.Set1(i)
			}
		}
	}
	return t
}
func findLMSPtrArray(t *bitvec.BitArr) []*subString {
	size := t.Size()
	p := make([]*subString, size)
	lastlms := -1
	for i := 1; i < size-1; i++ {
		if t.Get(i) == 1 && t.Get(i-1) == 0 {
			p[i] = &subString{start: uint32(i)}
			if lastlms != -1 {
				p[lastlms].end = uint32(i)
			}
			lastlms = i
		}
	}
	if lastlms != -1 {
		p[lastlms].end = uint32(size - 1)
	}
	p[size-1] = &subString{start: uint32(size - 1), end: uint32(size - 1)}
	return p
}
func putBucket(text []uint32, p []*subString, bucket []uint32, t *bitvec.BitArr) []int {
	bucketEnd := make([]uint32, len(bucket))
	copy(bucketEnd, bucket)
	sa := make([]int, len(text))
	n := len(sa)
	bucketHead := make([]uint32, len(bucket))
	bucketHead[0] = 0
	for i := 1; i < len(bucket); i++ {
		bucketHead[i] = bucket[i-1] + 1
	}
	//step 1
	for i := 0; i < n; i++ {
		sa[i] = -1
	}
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] != nil {
			sa[bucketEnd[text[p[i].start]]] = int(p[i].start)
			bucketEnd[text[p[i].start]]--
		}
	}
	//step 2
	for i := 0; i < n; i++ {
		if sa[i] != -1 {
			index := sa[i] - 1
			if index >= 0 && t.Get(index) == 0 {
				sa[bucketHead[text[index]]] = index
				bucketHead[text[index]]++
			}
		}
	}
	bucketEnd = make([]uint32, len(bucket))
	copy(bucketEnd, bucket)
	//step3
	for i := n - 1; i >= 0; i-- {
		if sa[i] != -1 {
			index := sa[i] - 1
			if index >= 0 && t.Get(index) == 1 {
				sa[bucketEnd[text[index]]] = index
				bucketEnd[text[index]]--
			}
		}
	}
	return sa
}

func compareSubstring(a subString, b subString, text []uint32, t *bitvec.BitArr) int {
	var size uint32 = 0
	if a.end-a.start+1 > b.end-b.start+1 {
		size = b.end - b.start + 1
	} else {
		size = a.end - a.start + 1
	}
	var i uint32 = 0
	for ; i < size; i++ {
		if text[a.start+i] == text[b.start+i] {
			atype := t.Get(int(a.start + i))
			btype := t.Get(int(b.start + i))
			if atype > btype {
				return 1
			} else if atype < btype {
				return -1
			}
		} else if text[a.start+i] < text[b.start+i] {
			return -1
		} else {
			return 1
		}
	}
	return 0
}
func nameP(p []*subString, sa []int, text []uint32, t *bitvec.BitArr) {
	n := len(sa)
	var j uint32 = 0
	lastlms := -1
	for i := 0; i < n; i++ {
		if p[sa[i]] != nil {
			if lastlms != -1 {
				result := compareSubstring(*p[lastlms], *p[sa[i]], text, t)
				if result < 0 {
					p[lastlms].name = j
					j++
				} else if result == 0 {
					p[lastlms].name = j
				} else if result > 0 {
					panic("bug")
				}
			}
			lastlms = sa[i]
		}
	}
	p[lastlms].name = j
}
func checkUnique(p []*subString) (bool, []int) {
	notNil := 0
	n := len(p)
	for i := 0; i < n; i++ {
		if p[i] != nil {
			notNil++
		}
	}
	sa := make([]int, notNil)
	for i := 0; i < notNil; i++ {
		sa[i] = -1
	}
	unique := true
	for i := 0; i < n; i++ {
		if p[i] != nil {
			if sa[p[i].name] == -1 {
				sa[p[i].name] = i
			} else {
				unique = false
				break
			}
		}
	}
	return unique, sa
}
