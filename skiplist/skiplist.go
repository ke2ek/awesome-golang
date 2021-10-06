package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_VALUE int = 1e9

type node struct {
	value int
	next  []*node
}

type Skiplist struct {
	level    int
	maxLevel int
	head     *node
	p        float64
}

func New(maxLv int, p float64) *Skiplist {
	sl := &Skiplist{level: 1, maxLevel: maxLv, p: p, head: &node{}}
	rand.Seed(time.Now().UnixNano())
	tail := &node{value: MAX_VALUE}
	for i := 0; i < sl.maxLevel; i++ {
		sl.head.next = append(sl.head.next, tail)
	}
	return sl
}

func (this *Skiplist) randomLevel() int {
	l := 1
	r := rand.Float64()
	for r < this.p && l < this.maxLevel {
		l++
		r = rand.Float64()
	}
	return l
}

// O(log(N))
func (this *Skiplist) Search(target int) bool {
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		for target > curr.next[i].value {
			curr = curr.next[i]
		}
	}
	curr = curr.next[0]
	return curr.value == target
}

// O(log(N))
func (this *Skiplist) Add(num int) {
	// Each node has integer as a value and pointers of next nodes as an array.
	// Start finding which location it is inserted at this.head.
	curr := this.head
	// Need to save pointers of the previous node in each level.
	prevNodes := make([]*node, this.maxLevel)
	// Find the insertion location
	for i := this.level - 1; i >= 0; i-- {
		// Go through the node having the value greater than a new number
		for num > curr.next[i].value {
			curr = curr.next[i]
		}
		// Save the pointer of the node which is lower bound of a new number in each level
		// by going down from the current level to 0
		prevNodes[i] = curr
	}
	// curr is gonna be what we want to find or upper bound of it at the bottom (a.k.a. 0)
	curr = curr.next[0]
	if curr.value == num {
		return
	}
	nextLevel := this.randomLevel()
	if nextLevel > this.level {
		// Fill empty pointers as what this.head is pointing
		for i := this.level; i < nextLevel; i++ {
			prevNodes[i] = this.head
		}
		this.level = nextLevel
	}
	// Create a new node
	curr = &node{value: num, next: make([]*node, this.maxLevel)}
	// Change next pointers of the new node.
	for i := 0; i < nextLevel; i++ {
		curr.next[i] = prevNodes[i].next[i]
		prevNodes[i].next[i] = curr
	}
}

// O(log(N))
func (this *Skiplist) Erase(num int) bool {
	curr := this.head
	prevNodes := make([]*node, this.maxLevel)
	for i := this.level - 1; i >= 0; i-- {
		for num > curr.next[i].value {
			curr = curr.next[i]
		}
		prevNodes[i] = curr
	}
	curr = curr.next[0]
	if curr.value != num {
		return false
	}
	for i := 0; i < this.level; i++ {
		if prevNodes[i].next[i] != curr {
			break
		}
		// Update next pointers of the previous node as them of the current node.
		prevNodes[i].next[i] = curr.next[i]
		curr.next[i] = nil
	}

	// If the next pointers of this.head include the pointer of tail node, then decrease this.level.
	for this.level > 1 && this.head.next[this.level-1].value == MAX_VALUE {
		this.level--
	}

	curr = nil
	return true
}

func (this *Skiplist) Print() {
	for i := 0; i < this.level; i++ {
		curr := this.head.next[i]
		fmt.Printf("Lv%d: H -> ", i+1)
		for curr.next != nil {
			fmt.Printf("%d -> ", curr.value)
			curr = curr.next[i]
		}
		fmt.Println("T")
	}
	fmt.Println()
}
