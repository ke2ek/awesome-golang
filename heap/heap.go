package heap

import "sort"

type H interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

func up(h H, i int) {
	for {
		j := (i - 1) / 2 // the index of parent
		// if the current location (=i) is the same as it of parent
		// or the value of the current location is not less than parent
		// then break
		if i == j || !h.Less(i, j) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}

func down(h H, from, to int) {
	here := from
	for {
		left, right := here*2+1, here*2+2
		if left >= to {
			break
		}
		next := here
		if h.Less(left, next) {
			next = left
		}
		if right < to && h.Less(right, next) {
			next = right
		}
		if here == next {
			break
		}
		h.Swap(here, next)
		here = next
	}
}

func Push(h H, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop(h H) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
