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

func (q Queue) Empty() bool {
	return q.back == -1
}

func (q Queue) Size() int {
	return q.back + 1
}

func (q Queue) Front() T {
	return q.elements[0].Value
}

func (q Queue) Back() T {
	return q.elements[q.back].Value
}

func (q *Queue) Push(value T) {
	if !q.Empty() && !common.EqualType(q.Back(), value) {
		panic("Error: type is not the same.")
	}
	q.elements = append(q.elements, &Element{value})
	q.back++
}

func (q *Queue) Pop() {
	q.elements = q.elements[1:]
	q.back--
}
