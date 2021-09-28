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

func (s Stack) Empty() bool {
	return s.back == -1
}

func (s Stack) Top() interface{} {
	return s.elements[s.back].Value
}

func (s Stack) Size() int {
	return s.back + 1
}

func (s *Stack) Push(value interface{}) {
	if !s.Empty() && !common.EqualType(s.Top(), value) {
		panic("Error: type is not the same.")
	}
	s.elements = append(s.elements, &Element{value})
	s.back++
}

func (s *Stack) Pop() {
	s.elements = s.elements[:s.back]
	s.back--
}
