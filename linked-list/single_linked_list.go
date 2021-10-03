package linkedlist

import "fmt"

type SingleLinkedList struct {
	head     *Node
	tail     *Node
	numNodes int
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (sll SingleLinkedList) Find(target T) *Node {
	node := sll.head
	for node != sll.tail {
		if node.Value == target {
			return node
		}
		node = node.next
	}
	return nil
}

func (sll SingleLinkedList) Empty() bool {
	return sll.numNodes == 0
}

func (sll SingleLinkedList) Size() int {
	return sll.numNodes
}

func (sll SingleLinkedList) Iterate() {
	node := sll.head
	for node != nil {
		fmt.Printf("%s ", node)
		node = node.next
	}
	fmt.Println()
}

func (sll SingleLinkedList) Front() T {
	return sll.head.Value
}

func (sll SingleLinkedList) Back() T {
	return sll.tail.Value
}

func (sll *SingleLinkedList) PushBack(target T) {
	node := Node{Value: target}
	if sll.tail == nil {
		sll.head = &node
		sll.tail = &node
	} else {
		sll.tail.next = &node
		sll.tail = &node
	}
	sll.numNodes++
}

func (sll *SingleLinkedList) PopBack() {
	if sll.Empty() {
		panic("PopBack(): list is empty.")
	}
	sll.numNodes--
	node := sll.head
	if node == sll.tail {
		sll.head = nil
		sll.tail = nil
		return
	}
	for node.next != sll.tail {
		node = node.next
	}
	node.next = nil
	sll.tail = node
}

func (sll *SingleLinkedList) PushFront(target T) {
	node := Node{Value: target, next: sll.head}
	sll.head = &node
	sll.numNodes++
	if sll.tail == nil {
		sll.tail = sll.head
	}
}

func (sll *SingleLinkedList) PopFront() {
	if sll.Empty() {
		panic("PopFront(): list is empty.")
	}
	sll.numNodes--
	newHead := sll.head.next
	if newHead == nil {
		sll.tail = nil
		return
	}
	sll.head.next = nil
	sll.head = newHead
}

func (sll *SingleLinkedList) Swap(l *SingleLinkedList) {
	h, t := sll.head, sll.tail
	sll.head, sll.tail = l.head, l.tail
	l.head, l.tail = h, t
}

func (sll *SingleLinkedList) Reverse() {
	sll.tail = sll.head
	prev, curr := sll.head, sll.head.next
	prev.next = nil
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev, curr = curr, next
	}
	sll.head = prev
}

func (sll *SingleLinkedList) Unique() {
	chk := map[string]*Node{}
	newsll := NewSingleLinkedList()
	node := sll.head
	for node != nil {
		str := fmt.Sprintf("%s", node.Value)
		if chk[str] == nil {
			chk[str] = node
			newsll.PushBack(node)
		} else {
			sll.numNodes--
		}
		node = node.next
	}
	sll.Swap(newsll)
}

/*
empty
size
front
back
insert -> double linked list
erase
push_back
push_front
pop_back
pop_front
swap : swaps the contents
reverse : reverses the order of the elements
unique : removes consecutive duplicate elements
merge : merges two sorted lists
iterate : visit
find
*/
