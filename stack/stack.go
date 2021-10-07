package stack

import "awesome-golang/common"

type Element common.Element

type Stack struct {
	elements []*Element
	back     int
}

func New() *Stack {
	return &Stack{back: -1}
}

func (this *Stack) Empty() bool {
	return this.back == -1
}

func (this *Stack) Top() interface{} {
	return this.elements[this.back].Value
}

func (this *Stack) Size() int {
	return this.back + 1
}

func (this *Stack) Push(value interface{}) {
	if !this.Empty() && !common.EqualType(this.Top(), value) {
		panic("Error: type is not the same.")
	}
	this.elements = append(this.elements, &Element{value})
	this.back++
}

func (this *Stack) Pop() {
	this.elements = this.elements[:this.back]
	this.back--
}
