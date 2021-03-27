package suffixArray

import (
	"4420project/bitvec"
	"4420project/waveletTree"
)

type tuple struct {
	key   string
	index int
}
type WTFMI struct {
	length  int
	charMap map[string]int
	c       []*tuple
	occ     *waveletTree.WaveletTree
}

func (this *WTFMI) OCC(str string, index int) int {
	if index < 0 {
		return 0
	}
	return this.occ.Rank(str, index)
}
func (this *WTFMI) C(char string) int {
	return this.c[this.charMap[char]].index
}
func (this *WTFMI) Search(pattern string) int {
	sp := 0
	ep := this.length - 1
	m := len(pattern)
	for i := m - 1; i >= 0; i-- {
		char := string(pattern[i])
		sp = this.C(char) + this.OCC(char, sp-1)
		ep = this.C(char) + this.OCC(char, ep) - 1
		if sp > ep {
			return 0
		}
	}
	return ep - sp + 1
}

type RLFMI struct {
	length  int
	charMap map[string]int
	c       []*tuple
	occ     *waveletTree.WaveletTree
	B       *bitvec.BasicBitVector
	B1      *bitvec.BasicBitVector
}

func (this *RLFMI) C(char string) int {
	return this.c[this.charMap[char]].index
}

func (this *RLFMI) OCC(c string, index int) int {
	if index < 0 {
		return 0
	}
	i := this.B.Rank1(index)
	j1 := this.occ.Rank(c, i-1)
	j := this.B1.Select1(this.C(c) + 1)
	equal := false
	if i > 1 && this.occ.Rank(c, i-1)-this.occ.Rank(c, i-2) == 1 {
		equal = true
	} else if i == 1 && j1 == 1 {
		equal = true
	}
	ofs := 0
	if equal {
		j1 = j1 - 1
		ofs = index - this.B.Select1(i) + 1
	}
	if this.C(c)+1+j1 > this.B1.Rank1(this.length-1) {
		return this.length - j + ofs
	}
	return this.B1.Select1(this.C(c)+1+j1) - j + ofs
}
func (this *RLFMI) Search(pattern string) int {
	sp := 0
	ep := this.length - 1
	m := len(pattern)
	for i := m - 1; i >= 0; i-- {
		char := string(pattern[i])
		j := this.B1.Select1(this.C(char) + 1)
		sp = j + this.OCC(char, sp-1)
		ep = j + this.OCC(char, ep) - 1
		if sp > ep {
			return 0
		}
	}
	return ep - sp + 1
}
