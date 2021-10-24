package tree

import "awesome-golang/common"

/*
Segment tree is used for storing information about intervals, or segments.
It represents each range of the given array as one node in a binary tree.
1. the root node always indicates the total range [0, n-1].
2. left and right subtree indicate the half of the range.
3. leaf node means the range having the one size (including only one element).
It is used to answer fastly questions about the specific range in an 1-D array.

e.g., Finding the minimum value of the specific range.
The query operation finds the minimum value by using the intersection of the current range and the required range.
*/

type RangeMinimumQuery struct {
	n        int
	rangeMin []int // It will stores the minimum value of each range.
}

func (this *RangeMinimumQuery) init(arr []int, left, right, node int) int {
	if left == right {
		this.rangeMin[node] = arr[left]
	} else {
		mid := (left + right) / 2
		leftMin := this.init(arr, left, mid, node*2)
		rightMin := this.init(arr, mid+1, right, node*2+1)
		this.rangeMin[node] = common.Min(leftMin, rightMin)
	}
	return this.rangeMin[node]
}

// O(log2(N))
func (this *RangeMinimumQuery) query(left, right, node, currentLeft, currentRight int) int {
	if currentLeft > right || currentRight < left { // out of range
		return 1e9
	}
	if currentLeft >= left && currentRight <= right {
		// If there is the current range inside the required range
		return this.rangeMin[node]
	}
	mid := (currentLeft + currentRight) / 2
	leftMin := this.query(left, right, node*2, currentLeft, mid)
	rightMin := this.query(left, right, node*2+1, mid+1, currentRight)
	return common.Min(leftMin, rightMin)
}

// Update a tree if array[index] is changed to newValue.
func (this *RangeMinimumQuery) update(index, value, node, currentLeft, currentRight int) int {
	if index < currentLeft || index > currentRight { // out of range
		return this.rangeMin[node]
	}
	if currentLeft == currentRight {
		this.rangeMin[node] = value
		return this.rangeMin[node]
	}
	mid := (currentLeft + currentRight) / 2
	leftMin := this.update(index, value, node*2, currentLeft, mid)
	rightMin := this.update(index, value, node*2+1, mid+1, currentRight)
	this.rangeMin[node] = common.Min(leftMin, rightMin)
	return this.rangeMin[node]
}

func (this *RangeMinimumQuery) Query(left, right int) int {
	return this.query(left, right, 1, 0, this.n-1)
}

func (this *RangeMinimumQuery) Update(index, value int) int {
	return this.update(index, value, 1, 0, this.n-1)
}

func NewRangeMinimumQuery(arr []int) *RangeMinimumQuery {
	tree := &RangeMinimumQuery{n: len(arr)}
	tree.rangeMin = make([]int, tree.n*4)
	tree.init(arr, 0, tree.n-1, 1)
	return tree
}
