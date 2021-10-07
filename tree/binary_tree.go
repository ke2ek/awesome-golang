package tree

import (
	"awesome-golang/queue"
)

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func NewNode(value interface{}, left *Node, right *Node) *Node {
	return &Node{Value: value, Left: left, Right: right}
}

// parent -> left -> right
func (this *Node) PreOrder(nodes *([]*Node)) {
	*nodes = append(*nodes, this)
	if this.Left != nil {
		this.Left.PreOrder(nodes)
	}
	if this.Right != nil {
		this.Right.PreOrder(nodes)
	}
}

// left -> parent -> right
func (this *Node) InOrder(nodes *([]*Node)) {
	if this.Left != nil {
		this.Left.InOrder(nodes)
	}
	*nodes = append(*nodes, this)
	if this.Right != nil {
		this.Right.InOrder(nodes)
	}
}

// left -> right -> parent
func (this *Node) PostOrder(nodes *([]*Node)) {
	if this.Left != nil {
		this.Left.PostOrder(nodes)
	}
	if this.Right != nil {
		this.Right.PostOrder(nodes)
	}
	*nodes = append(*nodes, this)
}

// Traverse level by level
func (this *Node) BFS(nodes *([][]*Node)) {
	q := queue.New()
	q.Push(this)
	for !q.Empty() {
		sz := q.Size()
		v := make([]*Node, sz)
		for i := 0; i < sz; i++ {
			here := q.Front().(*Node)
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
