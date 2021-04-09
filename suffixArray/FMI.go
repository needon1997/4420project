package suffixArray

import (
	"4420project/bitvec"
	"4420project/waveletTree"
)

type WTFMI struct {
	SASample  map[uint32]uint32
	InvSample map[uint32]uint32
	length    int
	charMap   map[byte]int
	c         []int
	occ       *waveletTree.WaveletTree
}

func (this *WTFMI) Substring(start int, end int) string {
	if end > this.length {
		panic("index out of bound")
	}
	l := end - start
	r := 0
	var a int
	for {
		a1, ok := this.InvSample[uint32(end)]
		if !ok {
			end = end + 1
			r++
		} else {
			a = int(a1)
			break
		}
	}
	for i := 0; i < r; i++ {
		a, _ = this.LF(a)
	}
	str := ""
	var char byte
	for i := 0; i < l; i++ {
		a, char = this.LF(a)
		str = string(char) + str
	}
	return str
}
func (this *WTFMI) Locate(index int) int {
	v := 0
	for {
		i, ok := this.SASample[uint32(index)]
		if !ok {
			index, _ = this.LF(index)
			v++
		} else {
			return int(i) + v
		}
	}
}

func (this *WTFMI) LF(i int) (int, byte) {
	char := this.occ.Get(i)
	return this.C(char) + this.OCC(char, i) - 1, char
}
func (this *WTFMI) OCC(str byte, index int) int {
	if index < 0 {
		return 0
	}
	return this.occ.Rank(str, index)
}
func (this *WTFMI) C(char byte) int {
	return this.c[this.charMap[char]]
}
func (this *WTFMI) Search(pattern string) int {
	sp := 0
	ep := this.length - 1
	m := len(pattern)
	for i := m - 1; i >= 0; i-- {
		char := pattern[i]
		sp = this.C(char) + this.OCC(char, sp-1)
		ep = this.C(char) + this.OCC(char, ep) - 1
		if sp > ep {
			return 0
		}
	}
	return ep - sp + 1
}

type RLFMI struct {
	SASample      map[uint32]uint32
	InvSample     map[uint32]uint32
	length        int
	charMap       map[byte]int
	c             []int
	distinctChars []byte
	occ           *waveletTree.WaveletTree
	B             *bitvec.BasicBitVector
	B1            *bitvec.BasicBitVector
	D             *bitvec.BasicBitVector
}

func (this *RLFMI) Substring(start int, end int) string {
	if end > this.length {
		panic("index out of bound")
	}
	l := end - start
	r := 0
	var a int
	for {
		a1, ok := this.InvSample[uint32(end)]
		if !ok {
			end = end + 1
			r++
		} else {
			a = int(a1)
			break
		}
	}
	for i := 0; i < r; i++ {
		a, _ = this.LF(a)
	}
	str := ""
	var char byte
	for i := 0; i < l; i++ {
		a, char = this.LF(a)
		str = string(char) + str
	}
	return str
}
func (this *RLFMI) Locate(index int) int {
	v := 0
	for {
		i, ok := this.SASample[uint32(index)]
		if !ok {
			index, _ = this.LF(index)
			v++
		} else {
			return int(i) + v
		}
	}
}
func (this *RLFMI) LF(i int) (int, byte) {
	i1 := this.B.Rank1(i)
	char := this.occ.Get(i1 - 1)
	return this.B1.Select1(this.C(char)+1) + this.OCC(char, i) - 1, char
}

func (this *RLFMI) C(char byte) int {
	return this.c[this.charMap[char]]
}

func (this *RLFMI) OCC(c byte, index int) int {
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
		char := pattern[i]
		j := this.B1.Select1(this.C(char) + 1)
		sp = j + this.OCC(char, sp-1)
		ep = j + this.OCC(char, ep) - 1
		if sp > ep {
			return 0
		}
	}
	return ep - sp + 1
}
