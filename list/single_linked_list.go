package list

import "fmt"

type SingleLinkedList struct {
	head     *Node
	tail     *Node
	numNodes int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (this *SingleLinkedList) Find(target T) *Node {
	node := this.head
	for node != this.tail {
		if node.Value == target {
			return node
		}
		node = node.next
	}
	return nil
}

func (this *SingleLinkedList) Empty() bool {
	return this.numNodes == 0
}

func (this *SingleLinkedList) Size() int {
	return this.numNodes
}

func (this *SingleLinkedList) Iterate() {
	node := this.head
	for node != nil {
		fmt.Printf("%s ", node)
		node = node.next
	}
	fmt.Println()
}

func (this *SingleLinkedList) First() *Node {
	return this.head
}

func (this *SingleLinkedList) Last() *Node {
	return this.tail
}

func (this *SingleLinkedList) PushBack(target T) {
	node := Node{Value: target}
	if this.tail == nil {
		this.head = &node
		this.tail = &node
	} else {
		this.tail.next = &node
		this.tail = &node
	}
	this.numNodes++
}

func (this *SingleLinkedList) PopBack() {
	if this.Empty() {
		panic("PopBack(): list is empty.")
	}
	this.numNodes--
	node := this.head
	if node == this.tail {
		this.head = nil
		this.tail = nil
		node.next = nil
		return
	}
	for node.next != this.tail {
		node = node.next
	}
	node.next = nil
	this.tail = node
}

func (this *SingleLinkedList) PushFront(target T) {
	node := Node{Value: target, next: this.head}
	this.head = &node
	this.numNodes++
	if this.tail == nil {
		this.tail = this.head
	}
}

func (this *SingleLinkedList) PopFront() {
	if this.Empty() {
		panic("PopFront(): list is empty.")
	}
	this.numNodes--
	newHead := this.head.next
	if newHead == nil {
		this.tail = nil
		return
	}
	this.head.next = nil
	this.head = newHead
}

func (this *SingleLinkedList) Swap(l *SingleLinkedList) {
	h, t := this.head, this.tail
	this.head, this.tail = l.head, l.tail
	l.head, l.tail = h, t
}

func (this *SingleLinkedList) Reverse() {
	this.tail = this.head
	prev, curr := this.head, this.head.next
	prev.next = nil
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev, curr = curr, next
	}
	this.head = prev
}

func (this *SingleLinkedList) Unique() {
	chk := map[string]*Node{}
	prev, node := this.head, this.head
	for node != nil {
		str := fmt.Sprintf("%s", node.Value)
		if chk[str] == nil {
			chk[str] = node
			prev, node = node, node.next
		} else {
			this.numNodes--
			removed := node
			node = node.next
			prev.next = removed.next
			removed.next = nil
			removed = nil
		}
	}
	this.tail = prev
}
