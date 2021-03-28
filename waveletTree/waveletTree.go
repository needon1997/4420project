package waveletTree

import (
	"4420project/bitvec"
)

type WaveletTree struct {
	root       *waveletNode
	mapping    map[string]int
	invMapping map[int]string
}

func (this *WaveletTree) Get(index int) string {
	encode := this.root.get(index)
	return this.invMapping[encode]
}
func (this *WaveletTree) Rank(char string, index int) int {
	if index < 0 {
		return 0
	}
	return this.root.rRank(this.mapping[char], index)
}
func NewWaveletTree(text string, chars []string) *WaveletTree {
	tree := &WaveletTree{root: nil, mapping: make(map[string]int, len(chars)), invMapping: make(map[int]string, len(chars))}
	for i := 0; i < len(chars); i++ {
		tree.mapping[chars[i]] = i
		tree.invMapping[i] = chars[i]
	}
	tree.root = newWaveletNode(text, tree.mapping)
	return tree
}
func newWaveletNode(text string, mapping map[string]int) *waveletNode {
	bitString := ""
	leftString := ""
	rightString := ""
	leftMapping := make(map[string]int, 1)
	rightMapping := make(map[string]int, 1)
	for i := 0; i < len(text); i++ {
		str := string(text[i])
		if mapping[str]%2 == 0 {
			bitString = bitString + "0"
			leftString += str
			leftMapping[str] = mapping[str] >> 1
		} else {
			bitString = bitString + "1"
			rightString += str
			rightMapping[str] = mapping[str] >> 1
		}
	}
	bitarr, err := bitvec.NewBitArr(bitString)
	if err != nil {
		panic(err)
	}
	bv := bitvec.NewBasicBitVec(bitarr)
	node := &waveletNode{binaryRank: bv}
	if len(leftMapping) <= 1 {
		node.leftChild = nil
	} else {
		node.leftChild = newWaveletNode(leftString, leftMapping)
	}
	if len(rightMapping) <= 1 {
		node.rightChild = nil
	} else {
		node.rightChild = newWaveletNode(rightString, rightMapping)
	}
	return node
}

type waveletNode struct {
	binaryRank *bitvec.BasicBitVector
	leftChild  *waveletNode
	rightChild *waveletNode
}

func (this *waveletNode) get(index int) int {
	encode := this.binaryRank.Get(index)
	if encode == 0 {
		rank := this.binaryRank.Rank0(index)
		if this.leftChild != nil {
			return int(encode) + (this.leftChild.get(rank-1) << 1)
		} else {
			return int(encode)
		}
	} else {
		rank := this.binaryRank.Rank1(index)
		if this.rightChild != nil {
			return int(encode) + (this.rightChild.get(rank-1) << 1)
		} else {
			return int(encode)
		}
	}
}
func (this *waveletNode) rRank(mapping int, index int) int {
	var rank int
	if mapping%2 == 0 {
		rank = this.binaryRank.Rank0(index)
		if this.leftChild != nil {
			return this.leftChild.rRank(mapping>>1, rank-1)
		} else {
			return rank
		}
	} else {
		rank = this.binaryRank.Rank1(index)
		if rank != 0 && this.rightChild != nil {
			return this.rightChild.rRank(mapping>>1, rank-1)
		} else {
			return rank
		}
	}
}
