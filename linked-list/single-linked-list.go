package linkedlist

type SingleLinkedList struct {
	head     *Node
	tail     *Node
	numNodes int
}

func New() *SingleLinkedList {
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

func (sll SingleLinkedList) Interate() {
	node := sll.head
	for node != sll.tail {
		node.Print()
		node = node.next
	}
}

func (sll SingleLinkedList) Front() T {
	return sll.head.Value
}

func (sll SingleLinkedList) Back() T {
	return sll.tail.Value
}

func (sll *SingleLinkedList) PushBack(target T) {
	sll.tail = &Node{Value: target}
	sll.numNodes++
}

func (sll *SingleLinkedList) PopBack() {
	sll.tail = nil
	sll.numNodes--
}

func (sll *SingleLinkedList) PushFront(target T) {
	newNode := Node{Value: target, next: sll.head}
	sll.head = &newNode
	sll.numNodes++
}

func (sll *SingleLinkedList) PopFront(target T) {
	newHead := sll.head.next
	sll.head.next = nil
	sll.head = newHead
	sll.numNodes--
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
	prev, curr := sll.head, sll.head.next
	for curr != nil {
		for prev.Value == curr.Value {
			removed := curr
			prev.next = curr.next
			curr = curr.next
			removed.next = nil
			removed = nil
		}
		prev, curr = curr, curr.next
	}
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
