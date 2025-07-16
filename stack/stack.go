package stack

import (
	"container/list"
	"errors"
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

func (s *Stack) Pop() (string, error) {
	if s.data.Len() == 0 {
		return "", errors.New("pop from empty stack")
	}
	elem := s.data.Back()
	s.data.Remove(elem)
	return elem.Value.(string), nil
}
