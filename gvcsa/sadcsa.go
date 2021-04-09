package gvcsa

import (
	"4420project/bitvec"
	"4420project/suffixArray"
	"math"
)

type SADCSArray struct {
	length      int
	A           []int
	Phi         []*EliasDeltaCoding
	B           []*bitvec.BasicBitVector
	D           *bitvec.BasicBitVector
	C           []byte
	invSASample []int
}

func (this *SADCSArray) Locate(index int) int {
	log2N := int(math.Ceil(math.Pow(math.Log2(float64(this.length)), 2)))
	i := index / log2N
	mod := index % log2N
	if mod == 0 {
		return this.invSASample[i]
	} else {
		index = index - mod
		sa := this.invSASample[i]
		for j := 0; j < mod; j++ {
			sa = this.Phi[0].Get(sa)
		}
		return sa
	}
}
func (this *SADCSArray) Substring(i, j int) string {
	length := j - i
	index := this.Locate(i)
	return this.subString(index, length, 0)
}
func (this *SADCSArray) Lookup(i int) int {
	index := this.rlookup(i, 0)
	if index < 0 {
		index = this.length + index
	}
	return index
}

func (this *SADCSArray) rlookup(i int, k int) int {
	if k == len(this.B) {
		return this.A[i]
	}
	var bi int
	if i == 0 {
		bi = this.B[k].Rank1(i)
	} else {
		bi = this.B[k].Rank1(i) - this.B[k].Rank1(i-1)
	}
	if bi == 1 {
		return 2 * this.rlookup(this.B[k].Rank1(i)-1, k+1)
	} else {
		if k == 0 {
			return this.rlookup(this.Phi[k].Get(i), k) - 1
		} else {
			return this.rlookup(this.Phi[k].Get(this.B[k].Rank0(i)-1), k) - 1
		}
	}

}
func (this *SADCSArray) subString(index int, length int, from int) string {
	for i := 0; i < from; i++ {
		index = this.Phi[0].Get(index)
	}
	str := make([]byte, length)
	for i := 0; i < length-from; i++ {
		str[i] = this.C[this.D.Rank1(index)-1]
		index = this.Phi[0].Get(index)
	}
	return string(str)
}
func (this *SADCSArray) Search(w string) int {
	Lw := this.getLw(w)
	Rw := this.getRw(w)
	sum := 0
	for i := Lw; i <= Rw; i++ {
		result, _ := this.compare(w, Lw, 0)
		if result == 0 {
			sum += 1 //fmt.Println(this.Text[this.Lookup(i) : this.Lookup(i)+len(w)])
		}
	}
	return sum
}
func (this *SADCSArray) getLw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := this.length
	result, l = this.compare(w, 0, 0)
	if result <= 0 {
		return 0
	}
	result, r = this.compare(w, n-1, 0)
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
		result, p := this.compare(w, M, h)
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

func (this *SADCSArray) getRw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := this.length
	result, l = this.compare(w, n-1, 0)
	if result >= 0 {
		return n
	}
	result, r = this.compare(w, 0, 0)
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
		result, p := this.compare(w, M, h)
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
func (this *SADCSArray) compare(a string, index int, from int) (result int, lcp int) {
	p := len(a)
	str := this.subString(index, p, from)
	for i := from; i < p; i++ {
		c := str[i-from]
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
func MakeSADCSArray(array suffixArray.SuffixArray) *SADCSArray {
	n := len(array.Text)
	csa := &SADCSArray{length: n}
	h := int(math.Log2(math.Log2(float64(n))))
	phi := make([]*EliasDeltaCoding, h)
	B := make([]*bitvec.BasicBitVector, h)
	a := preprocessSADRecur(array.Text, array.POS, 0, phi, B, h)

	D := bitvec.NewBitArrBySize(n)
	C := make([]byte, 0)
	D.Set1(0)
	C = append(C, array.Text[array.POS[0]])
	for i := 1; i < n; i++ {
		if array.Text[array.POS[i]] != array.Text[array.POS[i-1]] {
			D.Set1(i)
			C = append(C, array.Text[array.POS[i]])
		}
	}
	log2N := int(math.Ceil(math.Pow(math.Log2(float64(n)), 2)))
	invSASample := make([]int, int(math.Ceil(float64(n))/float64(log2N))+1)
	for i := 0; i < n; i++ {
		if array.POS[i]%log2N == 0 {
			invSASample[array.POS[i]/log2N] = i
		}
	}

	csa.A = a
	csa.B = B
	csa.Phi = phi
	csa.C = C
	csa.D = bitvec.NewBasicBitVec(D)
	csa.invSASample = invSASample
	return csa
}

func preprocessSADRecur(text string, sa []int, h int, phi []*EliasDeltaCoding, B []*bitvec.BasicBitVector, maxh int) []int {
	n := len(sa)
	a := make([]int, n-n/2)
	rank := make([]int, n+1)
	b := bitvec.NewBitArrBySize(n)
	j := 0
	//B A
	for i := 0; i < n; i++ {
		rank[sa[i]] = i
		if sa[i]%2 == 0 {
			b.Set1(i)
			a[j] = sa[i] / 2
			j++
		} else {
			b.Set0(i)
		}
	}
	rank[n] = rank[0]
	j = 0
	// Phi

	ph := make([]int, n/2)
	if h == 0 {
		ph = make([]int, n)
		for i := 0; i < n; i++ {
			ph[j] = rank[sa[i]+1]
			j++
		}
	} else {
		for i := 0; i < n; i++ {
			if b.Get(i) == 0 {
				ph[j] = rank[sa[i]+1]
				j++
			}
		}
	}

	phi[h] = NewEliasDeltaCoding(ph)
	B[h] = bitvec.NewBasicBitVec(b)
	if h == maxh-1 {
		return a
	} else {
		return preprocessRecur(text, a, h+1, phi, B, maxh)
	}
}
