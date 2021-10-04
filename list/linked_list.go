package list

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

func (this *Node) String() string {
	return fmt.Sprint(this.Value)
}

func (this *Node) Next() *Node {
	return this.next
}

func (this *Node) Prev() *Node {
	return this.prev
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
