package tree

import (
	"awesome-golang/common"
)

type RangeMinimumQuery struct {
	n        int
	arr      []int
	rangeMin []int // It will stores the minimum value of each range.
}

func (this *RangeMinimumQuery) outOfRange() interface{} {
	return int(1e9)
}

func (this *RangeMinimumQuery) setLeafNode(node, index int) {
	this.rangeMin[node] = this.arr[index]
}

func (this *RangeMinimumQuery) setNode(node int, value interface{}) {
	this.rangeMin[node] = value.(int)
}

func (this *RangeMinimumQuery) get(node int) interface{} {
	return this.rangeMin[node]
}

func (this *RangeMinimumQuery) merge(node int, leftResult, rightResult interface{}) interface{} {
	return common.Min(leftResult.(int), rightResult.(int))
}

func NewRangeMinimumQuery(arr []int) *RangeMinimumQuery {
	tree := &RangeMinimumQuery{
		n:        len(arr),
		arr:      arr,
		rangeMin: make([]int, 4*len(arr)),
	}
	Build(tree, 0, tree.n-1, 0)
	return tree
}

func (this *RangeMinimumQuery) Query(left, right int) int {
	return Query(this, left, right, 0, 0, this.n-1).(int)
}

func (this *RangeMinimumQuery) Update(index, value int) int {
	return Update(this, index, value, 0, 0, this.n-1).(int)
}
