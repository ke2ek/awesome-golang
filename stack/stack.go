package stack

import (
	"fmt"
	"reflect"
)

type Element struct {
	value interface{}
}

type Stack struct {
	elements []*Element
	back     int
}

func (e *Element) String() string {
	return fmt.Sprint(e.value)
}

func New() *Stack {
	return &Stack{back: -1}
}

func (s Stack) Empty() bool {
	return s.back == -1
}

func (s Stack) Top() interface{} {
	return s.elements[s.back].value
}

func (s Stack) Size() int {
	return s.back + 1
}

func (s *Stack) Push(value interface{}) {
	if !s.Empty() && reflect.TypeOf(s.Top()) != reflect.TypeOf(value) {
		panic("Error: type is not the same.")
	}
	s.elements = append(s.elements, &Element{value})
	s.back++
}

func (s *Stack) Pop() {
	s.elements = s.elements[:s.back]
	s.back--
}
