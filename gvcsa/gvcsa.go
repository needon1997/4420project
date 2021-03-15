package gvcsa

import (
	"4420project/bitvec"
	"4420project/suffixArray"
	"math"
)

type GVCSArray struct {
	Text string
	A    []int
	Phi  [][]int
	B    []*bitvec.BasicBitVector
}

func (this *GVCSArray) Lookup(i int) int {
	index := this.rlookup(i, 0)
	if index < 0{
		index = 219 + index + 1
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
		return this.rlookup(this.Phi[k][this.B[k].Rank0(i)-1], k) - 1
	}

}
func MakeGVCSArray(array suffixArray.SuffixArray) *GVCSArray {
	csa := &GVCSArray{Text: array.Text}
	h := int(math.Log2(math.Log2(float64(len(array.Text)))))
	phi := make([][]int, h)
	B := make([]*bitvec.BasicBitVector, h)
	a := preprocessRecur(array.POS, h, phi, B)
	csa.A = a
	csa.B = B
	csa.Phi = phi
	return csa
}

func preprocessRecur(sa []int, h int, phi [][]int, B []*bitvec.BasicBitVector) []int {
	n := len(sa)
	a := make([]int, n-n/2)
	ph := make([]int, n/2)
	rank := make([]int, n+1)
	b := bitvec.NewBitArrBySize(n)
	j := 0
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
	for i := 0; i < n; i++ {
		if b.Get(i) == 0 {
			ph[j] = rank[sa[i]+1]
			j++
		}
	}
	phi[len(phi)-h] = ph
	B[len(B)-h] = bitvec.NewBasicBitVec(b)

	if h == 1 {
		return a
	} else {
		return preprocessRecur(a, h-1, phi, B)
	}
}
