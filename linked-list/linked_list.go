package linkedlist

import "fmt"

type T interface{}

type Node struct {
	Value T
	next  *Node
	prev  *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type LinkedList interface {
	Find(target T) *Node
	Empty() bool
	Size() int
	Iterate()
	First() *Node
	Last() *Node
	PushBack()
	PopBack()
	PushFront()
	PopFront()
	Swap(l *LinkedList)
	Reverse()
	Unique()
}
