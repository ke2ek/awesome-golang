package linkedlist

import "fmt"

type T interface{}

type Node struct {
	Value T
	next  *Node
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type LinkedList interface {
	Find(target T) *Node
	Empty() bool
	Size() int
	Iterate()
	Front() T
	Back() T
	PushBack()
	PopBack()
	PushFront()
	PopFront()
	Swap(l *LinkedList)
	Reverse()
	Unique()
}
