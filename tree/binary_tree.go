package tree

import (
	"awesome-golang/queue"
)

type BTNode struct {
	Value interface{}
	Left  *BTNode
	Right *BTNode
}

// parent -> left -> right
func (this *BTNode) PreOrder(nodes *([]*BTNode)) {
	*nodes = append(*nodes, this)
	if this.Left != nil {
		this.Left.PreOrder(nodes)
	}
	if this.Right != nil {
		this.Right.PreOrder(nodes)
	}
}

// left -> parent -> right
func (this *BTNode) InOrder(nodes *([]*BTNode)) {
	if this.Left != nil {
		this.Left.InOrder(nodes)
	}
	*nodes = append(*nodes, this)
	if this.Right != nil {
		this.Right.InOrder(nodes)
	}
}

// left -> right -> parent
func (this *BTNode) PostOrder(nodes *([]*BTNode)) {
	if this.Left != nil {
		this.Left.PostOrder(nodes)
	}
	if this.Right != nil {
		this.Right.PostOrder(nodes)
	}
	*nodes = append(*nodes, this)
}

// Traverse level by level
func (this *BTNode) BFS(nodes *([][]*BTNode)) {
	q := queue.New()
	q.Push(this)
	for !q.Empty() {
		sz := q.Size()
		v := make([]*BTNode, sz)
		for i := 0; i < sz; i++ {
			here := q.Front().(*BTNode)
			v[i] = here
			q.Pop()
			if here.Left != nil {
				q.Push(here.Left)
			}
			if here.Right != nil {
				q.Push(here.Right)
			}
		}
		*nodes = append(*nodes, v)
	}
}
