package stack

import (
	"container/list"
)

type Stack struct {
	data *list.List
}

func New() *Stack {
	return &Stack{data: list.New()}
}

func (s *Stack) Push(value string) {
	s.data.PushBack(value)
}

func (s *Stack) Pop() string {
	if s.data.Len() == 0 {
		panic("can't pop from empty stack")

	}
	elem := s.data.Back()
	s.data.Remove(elem)
	return elem.Value.(string)
}

func (s *Stack) Empty() bool {
	return s.data.Len() == 0
}
