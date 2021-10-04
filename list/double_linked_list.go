package list

import "fmt"

type DoubleLinkedList struct {
	head     *Node
	tail     *Node
	numNodes int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (this *DoubleLinkedList) Find(target T) *Node {
	node := this.head
	for node != this.tail {
		if node.Value == target {
			return node
		}
		node = node.next
	}
	return nil
}

func (this *DoubleLinkedList) Empty() bool {
	return this.numNodes == 0
}

func (this *DoubleLinkedList) Size() int {
	return this.numNodes
}

func (this *DoubleLinkedList) Iterate() {
	node := this.head
	for node != nil {
		fmt.Printf("%s ", node)
		node = node.next
	}
	fmt.Println()
}

func (this *DoubleLinkedList) First() *Node {
	return this.head
}

func (this *DoubleLinkedList) Last() *Node {
	return this.tail
}

func (this *DoubleLinkedList) PushBack(target T) {
	node := Node{Value: target, prev: this.tail}
	if this.tail == nil {
		this.head = &node
		this.tail = &node
	} else {
		this.tail.next = &node
		this.tail = &node
	}
	this.numNodes++
}

func (this *DoubleLinkedList) PopBack() {
	if this.Empty() {
		panic("PopBack(): list is empty.")
	}
	this.numNodes--
	removed := this.tail
	this.tail = this.tail.prev
	this.tail.next = nil
	if removed == this.head {
		this.head = nil
		return
	}
	removed.prev, removed.next = nil, nil
	removed = nil
}

func (this *DoubleLinkedList) PushFront(target T) {
	node := Node{Value: target, next: this.head}
	this.head = &node
	this.numNodes++
	if this.tail == nil {
		this.tail = this.head
	}
}

func (this *DoubleLinkedList) PopFront() {
	if this.Empty() {
		panic("PopFront(): list is empty.")
	}
	this.numNodes--
	removed := this.head
	this.head = this.head.next
	this.head.prev = nil
	if removed == this.tail {
		this.tail = nil
		return
	}
	removed.prev, removed.next = nil, nil
	removed = nil
}

func (this *DoubleLinkedList) Swap(l *DoubleLinkedList) {
	h, t := this.head, this.tail
	this.head, this.tail = l.head, l.tail
	l.head, l.tail = h, t
}

func (this *DoubleLinkedList) Reverse() {
	l := NewDoubleLinkedList()
	for this.tail != nil {
		l.PushBack(this.tail.Value)
		this.tail = this.tail.prev
	}
	this.Swap(l)
}

func (this *DoubleLinkedList) Unique() {
	chk := map[string]*Node{}
	node := this.head
	for node != nil {
		str := fmt.Sprintf("%s", node.Value)
		if chk[str] == nil {
			chk[str] = node
			node = node.next
			this.tail = node
		} else {
			this.numNodes--
			removed := node
			this.tail = removed.prev
			node = node.next
			removed.prev.next = removed.next
			if removed.next != nil {
				removed.next.prev = removed.prev
			}
			removed.prev, removed.next = nil, nil
			removed = nil
		}
	}
}

func (this *DoubleLinkedList) Insert(to *Node, l *DoubleLinkedList) {
	if l.head == nil {
		return
	}
	if to == nil {
		this.tail.next = l.head
		l.head.prev = this.tail
		this.tail = l.tail
		return
	}
	if to.prev != nil {
		to.prev.next = l.head
		l.head.prev = to.prev
	} else {
		this.head = l.head
	}
	to.prev = l.tail
	l.tail.next = to
	this.numNodes += l.numNodes
}

func (this *DoubleLinkedList) Erase(target T) {
	removed := this.Find(target)
	if removed == nil {
		return
	}
	removed.prev.next = removed.next
	removed.next.prev = removed.prev
	removed.prev, removed.next = nil, nil
	removed = nil
	this.numNodes--
}
