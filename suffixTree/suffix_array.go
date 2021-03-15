package suffixTree

import (
	"fmt"
	"math"
)

type SuffixArray struct {
	Text string
	POS  []int
	LCP  []*LCPNode
	Llcp []int
	Rlcp []int
}

func (this *SuffixArray) initLCP() {
	//find the rank array
	rank := make([]int, len(this.POS))
	for i := 0; i < len(this.POS); i++ {
		rank[this.POS[i]] = i
	}
	h := 0
	for i := 0; i < len(this.POS); i++ {
		if rank[i] > 0 {
			j := this.POS[rank[i]-1]
			for this.Text[i+h] == this.Text[j+h] {
				h = h + 1
			}
			this.LCP[rank[i]] = &LCPNode{lcp: h, leaveCnt: 1}
			if h > 0 {
				h -= 1
			}
		}
	}
	validLcp := this.LCP[1:]
	_ = this.initLCPTreeRecur(&validLcp)
	//this.Llcp = make([]int, len(this.Text))
	//this.Rlcp = make([]int, len(this.Text))
	//this.buildLCPArray(root, 1, len(this.Text)-1) only available when the tree is balanced.
}

func (this *SuffixArray) initLCPTreeRecur(lcpArr *[]*LCPNode) *LCPNode {
	n := len(*lcpArr)
	if n == 2 {
		root := &LCPNode{lchild: (*lcpArr)[0], rchild: (*lcpArr)[1], lcp: int(math.Min(float64((*lcpArr)[0].lcp), float64((*lcpArr)[1].lcp)))}
		(*lcpArr)[0].parent = root
		(*lcpArr)[1].parent = root
		root.leaveCnt = (*lcpArr)[0].leaveCnt + (*lcpArr)[1].leaveCnt
		return root
	}
	newLcpArr := make([]*LCPNode, (n-1)/2+1)
	for i := 0; i < n; i += 2 {
		if i == n-1 {
			newNode := &LCPNode{lchild: (*lcpArr)[i], lcp: (*lcpArr)[i].lcp, leaveCnt: (*lcpArr)[i].leaveCnt}
			(*lcpArr)[i].parent = newNode
			newLcpArr[i/2] = newNode
			break
		}
		newNode := &LCPNode{lchild: (*lcpArr)[i], rchild: (*lcpArr)[i+1], lcp: int(math.Min(float64((*lcpArr)[i].lcp), float64((*lcpArr)[i+1].lcp)))}
		(*lcpArr)[i].parent = newNode
		(*lcpArr)[i+1].parent = newNode
		newNode.leaveCnt = (*lcpArr)[i].leaveCnt + (*lcpArr)[i+1].leaveCnt
		newLcpArr[i/2] = newNode
	}
	return this.initLCPTreeRecur(&newLcpArr)
}

//func (this *SuffixArray) buildLCPArray(node *LCPNode, start int, end int) {
//	if node.rchild == nil {
//		M := (start + end) / 2
//		this.Llcp[M] = node.lchild.lcp
//		return
//	}
//	if node.lchild.lchild == nil {
//		M := (start + end) / 2
//		this.Llcp[M] = node.lchild.lcp
//		this.Rlcp[M] = node.rchild.lcp
//		return
//	}
//	M := (start + end) / 2
//	this.Llcp[M] = node.lchild.lcp
//	this.Rlcp[M] = node.rchild.lcp
//	this.buildLCPArray(node.lchild, start, M)
//	this.buildLCPArray(node.rchild, M, end)
//}

func (this *SuffixArray) getLcp(index int, indey int) int {
	if index > indey {
		index = index + indey
		indey = index - indey
		index = index - indey
	}
	P := this.LCP[index].lcp
	Q := this.LCP[indey].lcp
	curX := this.LCP[index]
	curY := this.LCP[indey]
	meet := false
	var prev *LCPNode
	for !meet {
		if curX.parent == curY.parent {
			meet = true
			continue
		}
		prev = curX
		curX = curX.parent
		if curX.lchild == prev && P > curX.rchild.lcp {
			P = curX.rchild.lcp
		}
		prev = curY
		curY = curY.parent
		if curY.rchild == prev && Q > curY.lchild.lcp {
			Q = curX.rchild.lcp
		}

	}
	if P > Q {
		return Q
	} else {
		return P
	}

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

type LCPNode struct {
	parent   *LCPNode
	lchild   *LCPNode
	rchild   *LCPNode
	leaveCnt int
	lcp      int
}

func (this *SuffixArray) Search(w string) {
	Lw := this.getLw(w)
	Rw := this.getRw(w)
	for i := Lw; i <=
		Rw; i++ {
		result, _ := this.compare(w, this.POS[Lw], 0)
		if result == 0 {
			fmt.Println(this.Text[this.POS[i]:this.POS[i]+len(w)])
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
