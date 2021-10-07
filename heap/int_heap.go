package heap

type IntHeap []int

func NewIntHeap() *IntHeap {
	return &IntHeap{}
}

func (this *IntHeap) Len() int {
	return len(*this)
}

func (this *IntHeap) Less(i, j int) bool {
	return (*this)[i] < (*this)[j]
}

func (this *IntHeap) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *IntHeap) Pop() interface{} {
	last := this.Len() - 1
	x := (*this)[last]
	*this = (*this)[:last]
	return x
}
