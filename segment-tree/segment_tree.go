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

type SegmentTree interface {
	outOfRange() interface{}
	setLeafNode(node, index int)
	setNode(node int, value interface{})
	get(node int) interface{}
	merge(node int, leftResult, rightResult interface{}) interface{}
}

// O(log2(N))
func Build(segmentTree SegmentTree, left, right, node int) interface{} {
	if left == right {
		segmentTree.setLeafNode(node, left) // tree[node] = arr[left]
	} else {
		mid := (left + right) / 2
		leftResult := Build(segmentTree, left, mid, node*2+1)
		rightResult := Build(segmentTree, mid+1, right, node*2+2)
		mergedResult := segmentTree.merge(node, leftResult, rightResult)
		segmentTree.setNode(node, mergedResult)
	}
	return segmentTree.get(node)
}

// O(log2(N))
func Query(segmentTree SegmentTree, left, right, node, currentLeft, currentRight int) interface{} {
	if currentLeft > right || currentRight < left {
		return segmentTree.outOfRange()
	}
	if currentLeft >= left && currentRight <= right {
		// If there is the current range inside the required range
		return segmentTree.get(node)
	}
	mid := (currentLeft + currentRight) / 2
	leftResult := Query(segmentTree, left, right, node*2+1, currentLeft, mid)
	rightResult := Query(segmentTree, left, right, node*2+2, mid+1, currentRight)
	mergedResult := segmentTree.merge(node, leftResult, rightResult)
	return mergedResult
}

// O(log2(N))
// Update a tree if array[index] is changed to newValue.
func Update(segmentTree SegmentTree, index int, value interface{}, node, currentLeft, currentRight int) interface{} {
	if index < currentLeft || index > currentRight { // out of range
		return segmentTree.get(node)
	}
	if currentLeft == currentRight {
		segmentTree.setNode(node, value) // tree[node] = value
	} else {
		mid := (currentLeft + currentRight) / 2
		leftResult := Update(segmentTree, index, value, node*2+1, currentLeft, mid)
		rightResult := Update(segmentTree, index, value, node*2+2, mid+1, currentRight)
		mergedResult := segmentTree.merge(node, leftResult, rightResult)
		segmentTree.setNode(node, mergedResult)
	}
	return segmentTree.get(node)
}
