package waveletTree

import (
	"4420project/bitvec"
)

type WaveletTree struct {
	root       *waveletNode
	mapping    []byte
	invMapping map[int]byte
}

func (this *WaveletTree) Get(index int) byte {
	encode := this.root.get(index)
	return this.invMapping[encode]
}
func (this *WaveletTree) Rank(char byte, index int) int {
	if index < 0 || this.mapping[char] == 0 {
		return 0
	}
	return this.root.rRank(this.mapping[char], index)
}
func NewWaveletTree(text string, chars []byte) *WaveletTree {
	tree := &WaveletTree{root: nil, mapping: make([]byte, 256), invMapping: make(map[int]byte, len(chars))}
	for i := 0; i < len(chars); i++ {
		tree.mapping[chars[i]] = byte(i + 1)
		tree.invMapping[i+1] = chars[i]
	}
	tree.root = newWaveletNode([]byte(text), tree.mapping)
	return tree
}
func newWaveletNode(text []byte, mapping []byte) *waveletNode {
	bitarr := bitvec.NewBitArrBySize(len(text))
	leftString := make([]byte, 0)
	rightString := make([]byte, 0)
	leftMapping := make([]byte, 256)
	rightMapping := make([]byte, 256)
	leftTemp := make([]int, 256)
	rightTemp := make([]int, 256)
	leftCount := 0
	rightCount := 0
	for i := 0; i < 256; i++ {
		leftTemp[i] = -1
		rightTemp[i] = -1
	}
	for i := 0; i < len(text); i++ {
		str := text[i]
		if mapping[str]%2 == 0 {
			leftString = append(leftString, str)
			leftMapping[str] = mapping[str] >> 1
			if leftTemp[str] == -1 {
				leftCount += 1
				leftTemp[str] = 1
			}
		} else {
			bitarr.Set1(i)
			rightString = append(rightString, str)
			rightMapping[str] = mapping[str] >> 1
			if rightTemp[str] == -1 {
				rightCount += 1
				rightTemp[str] = 1
			}
		}
	}
	bv := bitvec.NewBasicBitVec(bitarr)
	node := &waveletNode{binaryRank: bv}
	if leftCount <= 1 {
		node.leftChild = nil
	} else {
		node.leftChild = newWaveletNode(leftString, leftMapping)
	}
	if rightCount <= 1 {
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
func (this *waveletNode) rRank(mapping byte, index int) int {
	var rank int
	if mapping%2 == 0 {
		rank = this.binaryRank.Rank0(index)
		if rank != 0 && this.leftChild != nil {
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
