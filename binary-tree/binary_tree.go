package tree

import (
	"awesome-golang/common"
	"awesome-golang/queue"
)

type I interface {
	Key() int
	Value() interface{}
	Left() interface{}
	Right() interface{}
}

// parent -> left -> right
func PreOrder(node I, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	*nodes = append(*nodes, node.Value())
	PreOrder(node.Left().(I), nodes)
	PreOrder(node.Right().(I), nodes)
}

// left -> parent -> right
func InOrder(node I, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	InOrder(node.Left().(I), nodes)
	*nodes = append(*nodes, node.Value())
	InOrder(node.Right().(I), nodes)
}

// left -> right -> parent
func PostOrder(node I, nodes *([]interface{})) {
	if common.IsNil(node) {
		return
	}
	PostOrder(node.Left().(I), nodes)
	PostOrder(node.Right().(I), nodes)
	*nodes = append(*nodes, node.Value())
}

// Traverse level by level
func BFS(root I) *([][]interface{}) {
	nodes := make([][]interface{}, 0)
	q := queue.New()
	q.Push(root)
	for !q.Empty() {
		sz := q.Size()
		v := make([]interface{}, sz)
		for i := 0; i < sz; i++ {
			here := q.Front().(I)
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
type Node struct {
	value interface{}
	left  *Node
	right *Node
}

func NewNode(value interface{}, left, right *Node) *Node {
	return &Node{value: value, left: left, right: right}
}

func (this *Node) Key() int           { return 1 } // not used
func (this *Node) Value() interface{} { return this.value }
func (this *Node) Left() interface{}  { return this.left }
func (this *Node) Right() interface{} { return this.right }
