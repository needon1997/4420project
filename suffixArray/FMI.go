package suffixArray

import "4420project/waveletTree"

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
