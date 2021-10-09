package tree

import (
	"awesome-golang/common"
	"awesome-golang/queue"
)

type Node interface {
	Key() interface{}
	Value() interface{}
	Left() interface{}
	Right() interface{}
}

// parent -> left -> right
func PreOrder(node Node, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	*nodes = append(*nodes, node.Value())
	PreOrder(node.Left().(Node), nodes)
	PreOrder(node.Right().(Node), nodes)
}

// left -> parent -> right
func InOrder(node Node, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	InOrder(node.Left().(Node), nodes)
	*nodes = append(*nodes, node.Value())
	InOrder(node.Right().(Node), nodes)
}

// left -> right -> parent
func PostOrder(node Node, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	PostOrder(node.Left().(Node), nodes)
	PostOrder(node.Right().(Node), nodes)
	*nodes = append(*nodes, node.Value())
}

// Traverse level by level
func BFS(root Node) *([][]interface{}) {
	nodes := make([][]interface{}, 0)
	q := queue.New()
	q.Push(root)
	for !q.Empty() {
		sz := q.Size()
		v := make([]interface{}, sz)
		for i := 0; i < sz; i++ {
			here := q.Front().(Node)
			v[i] = here.Value()
			q.Pop()
			if !common.IsNil(here.Left()) {
				q.Push(here.Left())
			}
			if !common.IsNil(here.Right()) {
				q.Push(here.Right())
			}
		}
		nodes = append(nodes, v)
	}
	return &nodes
}

// Implementation
type BTNode struct {
	value interface{}
	left  *BTNode
	right *BTNode
}

func NewBTNode(value interface{}, left, right *BTNode) *BTNode {
	return &BTNode{value: value, left: left, right: right}
}

func (this *BTNode) Key() interface{}   { return this } // not used
func (this *BTNode) Value() interface{} { return this.value }
func (this *BTNode) Left() interface{}  { return this.left }
func (this *BTNode) Right() interface{} { return this.right }
