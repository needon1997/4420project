package suffixArray

import (
	"4420project/bitvec"
	"4420project/waveletTree"
	"math"
)

type SuffixArray struct {
	Text string
	POS  []int
}

func (this *SuffixArray) BwtTransform() (string, string) {
	f := make([]byte, len(this.POS))
	l := make([]byte, len(this.POS))
	for i := 0; i < len(this.POS); i++ {
		f[i] = this.Text[this.POS[i]]
		if this.POS[i] > 0 {
			l[i] = this.Text[this.POS[i]-1]
		} else {
			l[i] = this.Text[len(this.Text)-1]
		}
	}
	return string(f), string(l)
}
func (this *SuffixArray) ToRLFMI() *RLFMI {
	_, tbwt := this.BwtTransform()
	S, B, B1, distinctChars, charMap, C := toRunLengthS(tbwt)
	occ := waveletTree.NewWaveletTree(S, distinctChars)
	n := len(this.Text)
	log2N := int(math.Ceil(math.Pow(math.Log2(float64(n)), 2)))
	//_ = int(math.Ceil(math.Log2(float64(n + 1))))
	SASample := make(map[int]int, int(math.Ceil(float64(n))/float64(log2N))+1)
	for i := 0; i < n; i++ {
		if this.POS[i]%log2N == 0 {
			SASample[i] = this.POS[i]
		}
	}
	InvSample := make(map[int]int, int(math.Ceil(float64(n))/float64(log2N))+1)
	for i := 0; i < n; i++ {
		if this.POS[i]%log2N == 0 && this.POS[i] > 0 {
			InvSample[this.POS[i]] = i
		}
		if this.POS[i] == n-1 {
			InvSample[n-1] = i
		}
	}
	return &RLFMI{length: len(this.Text), charMap: charMap, c: C, occ: occ, B: B, B1: B1, SASample: SASample, InvSample: InvSample, distinctChars: distinctChars}
}
func toRunLengthS(tbwt string) (string, *bitvec.BasicBitVector, *bitvec.BasicBitVector, []byte, map[byte]int, []int) {
	S := make([]byte, 0)
	length := len(tbwt)
	B := bitvec.NewBitArrBySize(length)
	lastChar := uint8(0)
	lastStart := -1
	distinctCharMap := make([]*[]int, 256)
	for i := 0; i < length; i++ {
		if distinctCharMap[tbwt[i]] == nil {
			distinctCharMap[tbwt[i]] = new([]int)
		}
		if tbwt[i] == lastChar {
			continue
		} else {
			if lastStart != -1 {
				*distinctCharMap[tbwt[lastStart]] = append(*distinctCharMap[tbwt[lastStart]], lastStart)
				*distinctCharMap[tbwt[lastStart]] = append(*distinctCharMap[tbwt[lastStart]], i-1)
			}
			B.Set1(i)
			S = append(S, tbwt[i])
			lastChar = tbwt[i]
			lastStart = i
		}
	}
	*distinctCharMap[tbwt[length-1]] = append(*distinctCharMap[tbwt[length-1]], lastStart)
	*distinctCharMap[tbwt[length-1]] = append(*distinctCharMap[tbwt[length-1]], length-1)
	B1 := bitvec.NewBitArrBySize(length)
	j := 0
	fS := make([]byte, 0)
	charMap := make(map[byte]int)
	distinctString := make([]byte, 0)
	m := 0
	for i := 0; i < 256; i++ {
		if distinctCharMap[i] != nil {
			distinctString = append(distinctString, byte(i))
			charMap[byte(i)] = m
			m++
			for k := 0; k < len(*distinctCharMap[i]); k += 2 {
				fS = append(fS, byte(i))
				B1.Set1(j)
				j++
				for l := (*distinctCharMap[i])[k] + 1; l <= (*distinctCharMap[i])[k+1]; l++ {
					B1.Set0(j)
					j++
				}
			}
		}
	}
	c := make([]int, 1)
	curChar := string(fS[0])
	c[0] = 0
	for i := 1; i < len(fS); i++ {
		if curChar == string(fS[i]) {
			continue
		} else {
			curChar = string(fS[i])
			c = append(c, i)
		}
	}
	return string(S), bitvec.NewBasicBitVec(B), bitvec.NewBasicBitVec(B1), distinctString, charMap, c
}
func (this *SuffixArray) ToWTFMI() *WTFMI {
	f, l := this.BwtTransform()
	c := make([]int, 0)
	distinctString := make([]byte, 0)
	curChar := f[0]
	c = append(c, 0)
	distinctString = append(distinctString, curChar)
	for i := 1; i < len(f); i++ {
		if curChar == f[i] {
			continue
		} else {
			curChar = f[i]
			c = append(c, i)
			distinctString = append(distinctString, curChar)
		}
	}
	charMap := make(map[byte]int)
	for i := 0; i < len(c); i++ {
		charMap[distinctString[i]] = i
	}
	occ := waveletTree.NewWaveletTree(l, distinctString)
	n := len(this.Text)
	log2N := int(math.Ceil(math.Pow(math.Log2(float64(n)), 2)))
	SASample := make(map[int]int, int(math.Ceil(float64(n))/float64(log2N))+1)
	InvSample := make(map[int]int, int(math.Ceil(float64(n))/float64(log2N))+1)
	for i := 0; i < n; i++ {
		if this.POS[i]%log2N == 0 {
			SASample[i] = this.POS[i]
			InvSample[this.POS[i]] = i

		}
		if this.POS[i] == n-1 {
			InvSample[n-1] = i
		}
	}
	return &WTFMI{charMap: charMap, c: c, occ: occ, length: len(this.Text), SASample: SASample, InvSample: InvSample}
}

func (this *SuffixArray) Search(w string) int {
	Lw := this.getLw(w)
	Rw := this.getRw(w)
	return Rw - Lw + 1
}
func (this *SuffixArray) getLw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := len(this.POS)
	result, l = this.compare(w, this.POS[0], 0)
	if result <= 0 {
		return 0
	}
	result, r = this.compare(w, this.POS[n-1], 0)
	if result > 0 {
		return n
	}
	L, R = 0, n-1
	for R-L > 1 {
		M := (L + R) / 2
		h := l
		if l > r {
			h = r
		}
		result, p := this.compare(w, this.POS[M], h)
		if result <= 0 {
			R = M
			r = p
		} else {
			L = M
			l = p
		}
	}
	return R
}

func (this *SuffixArray) getRw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := len(this.POS)
	result, l = this.compare(w, this.POS[n-1], 0)
	if result >= 0 {
		return n
	}
	result, r = this.compare(w, this.POS[0], 0)
	if result < 0 {
		return -1
	}
	L, R = 0, n-1
	for R-L > 1 {
		M := (L + R) / 2
		h := l
		if l > r {
			h = r
		}
		result, p := this.compare(w, this.POS[M], h)
		if result < 0 {
			R = M
			r = p
		} else {
			L = M
			l = p
		}
	}
	return L
}
func (this *SuffixArray) compare(a string, index int, from int) (result int, lcp int) {
	p := len(a)
	n := len(this.Text)
	for i := from; i < p; i++ {
		c := this.Text[int(math.Min(float64(index+i), float64(n-1)))]
		if a[i] == c {
			continue
		} else if a[i] < c {
			return -1, i
		} else if a[i] > c {
			return 1, i
		}
	}
	return 0, p
}
