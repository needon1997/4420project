package suffixArray

import (
	"4420project/waveletTree"
	"math"
)

type SuffixArray struct {
	Text string
	POS  []int
}

func (this *SuffixArray) BwtTransform() (string, string) {
	f := ""
	l := ""
	for i := 0; i < len(this.POS); i++ {
		f = f + string(this.Text[this.POS[i]])
		if this.POS[i] > 0 {
			l = l + string(this.Text[this.POS[i]-1])
		} else {
			l = l + string(this.Text[len(this.Text)-1])
		}
	}
	return f, l
}

func (this *SuffixArray) ToWTFMI() *WTFMI {
	f, l := this.BwtTransform()
	c := make([]*tuple, 1)
	curChar := string(f[0])
	c[0] = &tuple{key: curChar, index: 0}
	for i := 1; i < len(f); i++ {
		if curChar == string(f[i]) {
			continue
		} else {
			curChar = string(f[i])
			c = append(c, &tuple{curChar, i})
		}
	}
	distinctString := make([]string, len(c))
	charMap := make(map[string]int, len(c))
	for i := 0; i < len(c); i++ {
		distinctString[i] = c[i].key
		charMap[c[i].key] = i
	}
	occ := waveletTree.NewWaveletTree(l, distinctString)
	return &WTFMI{charMap: charMap, c: c, occ: occ, length: len(this.Text)}
}

func (this *SuffixArray) Search(w string) {
	Lw := this.getLw(w)
	Rw := this.getRw(w)
	for i := Lw; i <=
		Rw; i++ {
		result, _ := this.compare(w, this.POS[Lw], 0)
		if result == 0 {
			//fmt.Println(this.Text[this.POS[i] : this.POS[i]+len(w)])
		}
	}
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
