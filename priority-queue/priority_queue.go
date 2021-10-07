package priorityqueue

import "awesome-golang/heap"

type T interface{}

type Node struct {
	Key   int
	Value interface{}
}

type Heapq []*Node

func (this *Heapq) Len() int {
	return len(*this)
}

func (this *Heapq) Less(i, j int) bool {
	return (*this)[i].Key < (*this)[j].Key
}

func (this *Heapq) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *Heapq) Push(x interface{}) {
	node, ok := x.(*Node)
	if !ok {
		panic("Push(): the type of a given parameter is not the pointer type of Node.")
	}
	*this = append(*this, node)
}

func (this *Heapq) Pop() interface{} {
	last := this.Len() - 1
	x := (*this)[last]
	*this = (*this)[:last]
	return x
}

type PriorityQueue struct {
	q *Heapq
}

func New() *PriorityQueue {
	return &PriorityQueue{q: &Heapq{}}
}

func (this *PriorityQueue) Push(key int, value interface{}) {
	heap.Push(this.q, &Node{key, value})
}

func (this *PriorityQueue) Pop() *Node {
	return heap.Pop(this.q).(*Node)
}

func (this *PriorityQueue) Size() int {
	return this.q.Len()
}

func (this *PriorityQueue) Empty() bool {
	return this.q.Len() == 0
}
