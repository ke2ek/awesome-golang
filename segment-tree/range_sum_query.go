package tree

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

type RangeSumQuery struct {
	n        int
	rangeSum []int // It will stores the minimum value of each range.
}

func (this *RangeSumQuery) init(arr []int, left, right, node int) int {
	if left == right {
		this.rangeSum[node] = arr[left]
	} else {
		mid := (left + right) / 2
		leftSum := this.init(arr, left, mid, node*2)
		rightSum := this.init(arr, mid+1, right, node*2+1)
		this.rangeSum[node] = leftSum + rightSum
	}
	return this.rangeSum[node]
}

// O(log2(N))
func (this *RangeSumQuery) query(left, right, node, currentLeft, currentRight int) int {
	if currentLeft > right || currentRight < left { // out of range
		return 0
	}
	if left <= currentLeft && right >= currentRight {
		// If there is the current range inside the required range
		return this.rangeSum[node]
	}
	mid := (currentLeft + currentRight) / 2
	leftSum := this.query(left, right, node*2, currentLeft, mid)
	rightSum := this.query(left, right, node*2+1, mid+1, currentRight)
	return leftSum + rightSum
}

// Update a tree if array[index] is changed to newValue.
func (this *RangeSumQuery) update(index, value, node, currentLeft, currentRight int) int {
	if index < currentLeft || index > currentRight { // out of range
		return this.rangeSum[node]
	}
	if currentLeft == currentRight {
		this.rangeSum[node] = value
		return this.rangeSum[node]
	}
	mid := (currentLeft + currentRight) / 2
	leftSum := this.update(index, value, node*2, currentLeft, mid)
	rightSum := this.update(index, value, node*2+1, mid+1, currentRight)
	this.rangeSum[node] = leftSum + rightSum
	return this.rangeSum[node]
}

func (this *RangeSumQuery) Query(left, right int) int {
	return this.query(left, right, 1, 0, this.n-1)
}

func (this *RangeSumQuery) Update(index, value int) int {
	return this.update(index, value, 1, 0, this.n-1)
}

func NewRangeSumQuery(arr []int) *RangeSumQuery {
	tree := &RangeSumQuery{n: len(arr)}
	tree.rangeSum = make([]int, tree.n*4)
	tree.init(arr, 0, tree.n-1, 1)
	return tree
}
