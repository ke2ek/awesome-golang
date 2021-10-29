package tree

type RangeSumQuery struct {
	n        int
	arr      []int
	rangeSum []int // It will stores the minimum value of each range.
}

func (this *RangeSumQuery) outOfRange() interface{} {
	return 0
}

func (this *RangeSumQuery) setLeafNode(node, index int) {
	this.rangeSum[node] = this.arr[index]
}

func (this *RangeSumQuery) setNode(node int, value interface{}) {
	this.rangeSum[node] = value.(int)
}

func (this *RangeSumQuery) get(node int) interface{} {
	return this.rangeSum[node]
}

func (this *RangeSumQuery) merge(node int, leftResult, rightResult interface{}) interface{} {
	return leftResult.(int) + rightResult.(int)
}

func NewRangeSumQuery(arr []int) *RangeSumQuery {
	tree := &RangeSumQuery{
		n:        len(arr),
		arr:      arr,
		rangeSum: make([]int, 4*len(arr)),
	}
	Build(tree, 0, tree.n-1, 0)
	return tree
}

func (this *RangeSumQuery) Query(left, right int) int {
	return Query(this, left, right, 0, 0, this.n-1).(int)
}

func (this *RangeSumQuery) Update(index, value int) int {
	return Update(this, index, value, 0, 0, this.n-1).(int)
}
