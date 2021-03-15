package suffixTree

type Edge struct {
	Start int
	End   *int
	From  *Node
	To    *Node
	Index int
}

type Node struct {
	Edges      []*Edge
	SuffixNode *Node
	EdgeNum    byte
}

func NewNode() *Node {
	node := &Node{Edges: make([]*Edge, 256)}
	node.SuffixNode = node
	return node
}

type SuffixTree struct {
	Root *Node
	Text string
}

func NewSuffixTree(text string) *SuffixTree {
	sTree := &SuffixTree{Text: text, Root: NewNode()}
	sTree.build()
	return sTree
}
func (this *SuffixTree) build() {
	var (
		end          *int = new(int)
		remainSuffix *int = new(int)
		activeNode   **Node
		activeEdge   *int = new(int)
		activeLength *int = new(int)
		index        int
		textLength   int
	)
	*end = -1
	*remainSuffix = 0
	activeNode = new(*Node)
	*activeNode = this.Root
	*activeEdge = -1
	*activeLength = 0
	index = 0
	this.Text = this.Text + "$"
	textLength = len(this.Text)
	for index < textLength {
		this.extend(index, end, remainSuffix, activeEdge, activeLength, activeNode)
		index = index + 1
	}
	var init_index *int = new(int)
	*init_index = -1
	this.setSuffixIndex(this.Root, init_index)
}

func (this *SuffixTree) setSuffixIndex(node *Node, index *int) {
	for _, val := range node.Edges {
		if val == nil {
			continue
		}
		if val.To != nil {
			val.Index = -1
			this.setSuffixIndex(val.To, index)
		} else {
			*index = *index + 1
			val.Index = *index
		}
	}
}

func (this *SuffixTree) ToSuffixTray() {
	return
}

func (this *SuffixTree) ToSuffixArray() *SuffixArray {
	index := new(int)
	*index = 0
	suffixArray := &SuffixArray{POS: make([]int, len(this.Text)), LCP: make([]*LCPNode, len(this.Text)), Text: this.Text}
	pos := make([]int, len(this.Text))
	suffixArray.POS = pos
	this.toSuffixArrayRecur(this.Root, "", index, &pos)
	suffixArray.initLCP()
	return suffixArray
}

func (this *SuffixTree) toSuffixArrayRecur(node *Node, prefix string, index *int, array *[]int) {
	for _, edge := range node.Edges {
		if edge == nil {
			continue
		}
		cur_prefix := prefix + this.Text[edge.Start:*(edge.End)+1]
		if edge.To != nil {
			this.toSuffixArrayRecur(edge.To, cur_prefix, index, array)
		} else {
			(*array)[*index] = len(this.Text) - len(cur_prefix)
			*index = *index + 1
		}
	}
}

func (this *SuffixTree) extend(index int, end, remainSuffix, activeEdge, activeLength *int, activeNode **Node) {
	var edge *Edge
	var edgeLength int
	var lastNode *Node = nil
	*remainSuffix = *remainSuffix + 1
	*end = *end + 1
	for *remainSuffix > 0 {
		if *activeLength == 0 {
			*activeEdge = index
		}
		edge = this.getEdge(*activeEdge, activeNode)
		if edge != nil {
			edgeLength = *edge.End - edge.Start + 1
			if this.walkDown(edgeLength, activeLength, activeEdge, edge, activeNode) {
				continue
			}
			//rule 3
			if this.Text[index] == this.Text[edge.Start+*activeLength] {
				*activeLength = *activeLength + 1
				if lastNode != nil && *activeNode != this.Root {
					lastNode.SuffixNode = *activeNode
					lastNode = nil
				}
				break
			} else { //rule 2, split the edge and create a new node
				newNode := &Node{SuffixNode: this.Root, Edges: make([]*Edge, 256)}
				newEdge1 := &Edge{Start: edge.Start + *activeLength, To: edge.To, From: newNode, Index: -1}
				if edge.End == end {
					newEdge1.End = end
				} else {
					newEdge1.End = new(int)
					*newEdge1.End = *edge.End
				}
				edge.End = new(int)
				*edge.End = edge.Start + *activeLength - 1
				edge.To = newNode
				newNode.Edges[this.Text[newEdge1.Start]] = newEdge1
				newNode.EdgeNum++
				newEdge2 := &Edge{Start: index, End: end, From: newNode, To: nil, Index: -1}
				newNode.EdgeNum++
				newNode.Edges[this.Text[newEdge2.Start]] = newEdge2
				if lastNode != nil {
					lastNode.SuffixNode = newNode
				}
				lastNode = newNode
			}

		} else {
			edge = &Edge{Start: index, End: end, From: *activeNode, To: nil, Index: -1}
			(*activeNode).Edges[this.Text[edge.Start]] = edge
			(*activeNode).EdgeNum++
			if lastNode != nil {
				lastNode.SuffixNode = *activeNode
				lastNode = nil
			}
		}
		*remainSuffix = *remainSuffix - 1
		if *activeNode == this.Root && *activeLength > 0 {
			*activeEdge = index - *remainSuffix + 1
			*activeLength = *activeLength - 1
		} else {
			*activeNode = (*activeNode).SuffixNode
		}
	}
}

func (this *SuffixTree) getEdge(activeEdge int, activeNode **Node) (result *Edge) {
	result = (*activeNode).Edges[this.Text[activeEdge]]
	return
}

func (this *SuffixTree) walkDown(edgeLength int, activeLength, activeEdge *int, edge *Edge, activeNode **Node) bool {
	if *activeLength >= edgeLength {
		*activeNode = edge.To
		*activeEdge = *activeEdge + edgeLength
		*activeLength = *activeLength - edgeLength
		return true
	}
	return false
}
