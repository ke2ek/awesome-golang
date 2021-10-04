package linkedlist

import "fmt"

type DoubleLinkedList struct {
	head     *Node
	tail     *Node
	numNodes int
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (dll DoubleLinkedList) Find(target T) *Node {
	node := dll.head
	for node != dll.tail {
		if node.Value == target {
			return node
		}
		node = node.next
	}
	return nil
}

func (dll DoubleLinkedList) Empty() bool {
	return dll.numNodes == 0
}

func (dll DoubleLinkedList) Size() int {
	return dll.numNodes
}

func (dll DoubleLinkedList) Iterate() {
	node := dll.head
	for node != nil {
		fmt.Printf("%s ", node)
		node = node.next
	}
	fmt.Println()
}

func (dll DoubleLinkedList) Front() T {
	return dll.head.Value
}

func (dll DoubleLinkedList) Back() T {
	return dll.tail.Value
}

func (dll *DoubleLinkedList) PushBack(target T) {
	node := Node{Value: target, prev: dll.tail}
	if dll.tail == nil {
		dll.head = &node
		dll.tail = &node
	} else {
		dll.tail.next = &node
		dll.tail = &node
	}
	dll.numNodes++
}

func (dll *DoubleLinkedList) PopBack() {
	if dll.Empty() {
		panic("PopBack(): list is empty.")
	}
	dll.numNodes--
	removed := dll.tail
	dll.tail = dll.tail.prev
	dll.tail.next = nil
	if removed == dll.head {
		dll.head = nil
		return
	}
	removed.prev, removed.next = nil, nil
	removed = nil
}

func (dll *DoubleLinkedList) PushFront(target T) {
	node := Node{Value: target, next: dll.head}
	dll.head = &node
	dll.numNodes++
	if dll.tail == nil {
		dll.tail = dll.head
	}
}

func (dll *DoubleLinkedList) PopFront() {
	if dll.Empty() {
		panic("PopFront(): list is empty.")
	}
	dll.numNodes--
	removed := dll.head
	dll.head = dll.head.next
	dll.head.prev = nil
	if removed == dll.tail {
		dll.tail = nil
		return
	}
	removed.prev, removed.next = nil, nil
	removed = nil
}

func (dll *DoubleLinkedList) Swap(l *DoubleLinkedList) {
	h, t := dll.head, dll.tail
	dll.head, dll.tail = l.head, l.tail
	l.head, l.tail = h, t
}

func (dll *DoubleLinkedList) Reverse() {
	l := NewDoubleLinkedList()
	for dll.tail != nil {
		l.PushBack(dll.tail.Value)
		dll.tail = dll.tail.prev
	}
	dll.Swap(l)
}

func (dll *DoubleLinkedList) Unique() {
	chk := map[string]*Node{}
	node := dll.head
	for node != nil {
		str := fmt.Sprintf("%s", node.Value)
		if chk[str] == nil {
			chk[str] = node
			node = node.next
			dll.tail = node
		} else {
			dll.numNodes--
			removed := node
			dll.tail = removed.prev
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

func (dll *DoubleLinkedList) Insert(to *Node, l *DoubleLinkedList) {
	if l.head == nil {
		return
	}
	if to == nil {
		dll.tail.next = l.head
		l.head.prev = dll.tail
		dll.tail = l.tail
		return
	}
	if to.prev != nil {
		to.prev.next = l.head
		l.head.prev = to.prev
	} else {
		dll.head = l.head
	}
	to.prev = l.tail
	l.tail.next = to
	dll.numNodes += l.numNodes
}

func (dll *DoubleLinkedList) Erase(target T) {
	removed := dll.Find(target)
	if removed == nil {
		return
	}
	removed.prev.next = removed.next
	removed.next.prev = removed.prev
	removed.prev, removed.next = nil, nil
	removed = nil
	dll.numNodes--
}

// if it returns -1, lhs is less than rhs.
// if it returns 1, lhs is greater than rhs.
// if it returns 0, lhs equals to rhs.
type fn func(T, T) int

// Assume that two lists have been sorted already.
func Merge(l1 *DoubleLinkedList, l2 *DoubleLinkedList, op fn) *DoubleLinkedList {
	dll := NewDoubleLinkedList()
	n1, n2 := l1.head, l2.head
	for n1 != l1.tail && n2 != l2.tail {
		if op(n1.Value, n2.Value) <= 0 {
			dll.PushBack(n1.Value)
			n1 = n1.next
		} else {
			dll.PushBack(n2.Value)
			n2 = n2.next
		}
	}
	return nil
}
