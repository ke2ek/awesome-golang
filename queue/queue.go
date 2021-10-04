package queue

import "awesome-golang/common"

type T interface{}
type Element common.Element

type Queue struct {
	elements []*Element
	back     int
}

func New() *Queue {
	return &Queue{back: -1}
}

func (this *Queue) Empty() bool {
	return this.back == -1
}

func (this *Queue) Size() int {
	return this.back + 1
}

func (this *Queue) Front() T {
	return this.elements[0].Value
}

func (this *Queue) Back() T {
	return this.elements[this.back].Value
}

func (this *Queue) Push(value T) {
	if !this.Empty() && !common.EqualType(this.Back(), value) {
		panic("Error: type is not the same.")
	}
	this.elements = append(this.elements, &Element{value})
	this.back++
}

func (this *Queue) Pop() {
	this.elements = this.elements[1:]
	this.back--
}
