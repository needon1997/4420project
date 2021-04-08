package waveletTree

import (
	"4420project/bitvec"
)

type WaveletTree2 struct {
	mapping    map[byte]int
	invMapping map[int]byte
	nodes      []bitvec.BasicBitVector
	treemap    *bitvec.BasicBitVector
}

func (this *WaveletTree2) Rank(char byte, index int) int {
	if index < 0 {
		return 0
	}
	return this.rRank(1, this.mapping[char], index)
}
func (this *WaveletTree2) rRank(i int, mapping int, index int) int {
	var rank int
	lc := 2 * this.treemap.Rank1(i)
	if mapping%2 == 0 {
		rank = this.nodes[i-1].Rank0(index)
		if rank != 0 && this.nodes[lc-1].Size() != 0 {
			return this.rRank(lc, mapping>>1, rank-1)
		} else {
			return rank
		}
	} else {
		rank = this.nodes[i-1].Rank1(index)
		if rank != 0 && this.nodes[lc-1].Size() != 0 {
			return this.rRank(lc+1, mapping>>1, rank-1)
		} else {
			return rank
		}
	}
}
func NewWaveletTree2(text string, chars []byte) *WaveletTree2 {
	tree := &WaveletTree2{mapping: make(map[byte]int, len(chars)), invMapping: make(map[int]byte, len(chars))}
	for i := 0; i < len(chars); i++ {
		tree.mapping[chars[i]] = i
		tree.invMapping[i] = chars[i]
	}
	tree.nodes = make([]bitvec.BasicBitVector, 2*len(chars)+1)
	b := bitvec.NewBitArrBySize(2*len(chars) + 2)
	tree.initializeNode(1, []byte(text), tree.mapping, b)
	tree.treemap = bitvec.NewBasicBitVec(b)
	return tree
}

func (this *WaveletTree2) initializeNode(index int, text []byte, mapping map[byte]int, b *bitvec.BitArr) {
	b.Set1(index)
	bitarr := bitvec.NewBitArrBySize(len(text))
	leftString := make([]byte, 0)
	rightString := make([]byte, 0)
	leftMapping := make(map[byte]int, 1)
	rightMapping := make(map[byte]int, 1)
	for i := 0; i < len(text); i++ {
		str := text[i]
		if mapping[str]%2 == 0 {
			leftString = append(leftString, str)
			leftMapping[str] = mapping[str] >> 1
		} else {
			bitarr.Set1(i)
			rightString = append(rightString, str)
			rightMapping[str] = mapping[str] >> 1
		}
	}
	this.nodes[index-1] = *(bitvec.NewBasicBitVec(bitarr))
	if len(leftMapping) > 1 {
		this.initializeNode(2*b.Rank1(index), leftString, leftMapping, b)
	}
	if len(rightMapping) > 1 {
		this.initializeNode(2*b.Rank1(index)+1, rightString, rightMapping, b)
	}
}
