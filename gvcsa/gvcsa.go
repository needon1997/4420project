package gvcsa

import (
	"4420project/bitvec"
	"4420project/suffixArray"
	"math"
)

type GVCSArray struct {
	Text string
	A    []int
	Phi  []*EliasDeltaCoding
	B    []*bitvec.BasicBitVector
}

func (this *GVCSArray) Lookup(i int) int {
	index := this.rlookup(i, 0)
	if index < 0 {
		index = len(this.Text) + index
	}
	return index
}

func (this *GVCSArray) rlookup(i int, k int) int {
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
		return this.rlookup(this.Phi[k].Get(this.B[k].Rank0(i)-1), k) - 1
	}

}

func (this *GVCSArray) Search(w string) int {
	Lw := this.getLw(w)
	Rw := this.getRw(w)
	sum := 0
	for i := Lw; i <= Rw; i++ {
		result, _ := this.compare(w, this.Lookup(Lw), 0)
		if result == 0 {
			sum += 1 //fmt.Println(this.Text[this.Lookup(i) : this.Lookup(i)+len(w)])
		}
	}
	return sum
}
func (this *GVCSArray) getLw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := len(this.Text)
	result, l = this.compare(w, this.Lookup(0), 0)
	if result <= 0 {
		return 0
	}
	result, r = this.compare(w, this.Lookup(n-1), 0)
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
		result, p := this.compare(w, this.Lookup(M), h)
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

func (this *GVCSArray) getRw(w string) int {
	var L int
	var R int
	var l int
	var r int
	var result int
	n := len(this.Text)
	result, l = this.compare(w, this.Lookup(n-1), 0)
	if result >= 0 {
		return n
	}
	result, r = this.compare(w, this.Lookup(0), 0)
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
		result, p := this.compare(w, this.Lookup(M), h)
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
func (this *GVCSArray) compare(a string, index int, from int) (result int, lcp int) {
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

func MakeGVCSArray(array suffixArray.SuffixArray) *GVCSArray {
	csa := &GVCSArray{Text: array.Text}
	h := int(math.Log2(math.Log2(float64(len(array.Text)))))
	phi := make([]*EliasDeltaCoding, h)
	B := make([]*bitvec.BasicBitVector, h)
	a := preprocessRecur(array.Text, array.POS, 0, phi, B, h)

	csa.A = a
	csa.B = B
	csa.Phi = phi
	return csa
}

func preprocessRecur(text string, sa []int, h int, phi []*EliasDeltaCoding, B []*bitvec.BasicBitVector, maxh int) []int {
	n := len(sa)
	a := make([]int, n-n/2)
	ph := make([]int, n/2)
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
	for i := 0; i < n; i++ {
		if b.Get(i) == 0 {
			ph[j] = rank[sa[i]+1]
			j++
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
