package datastructures

import "errors"

type ListStack[T any] struct {
	length int
	top    *node[T]
}

func NewListStack[T any]() ListStack[T] {
	return ListStack[T]{}
}

func (s *ListStack[T]) Push(value T) {
	newNode := &node[T]{}
	newNode.value = value

	if s.IsEmpty() {
		s.top = newNode
		s.length = 1
	} else {
		oldTop := s.top
		newNode.previous = oldTop
		s.top = newNode
		s.length++
	}
}

func (s *ListStack[T]) Pop() (popped T, err error) {
	if s.IsEmpty() {
		return popped, errors.New("stack is empty")
	}

	popped = s.top.value
	s.top = s.top.previous
	s.length--

	return popped, nil
}

func (s ListStack[T]) Peek() (value T, err error) {
	if s.IsEmpty() {
		return value, errors.New("stack is empty")
	}

	return s.top.value, nil
}

func (s ListStack[T]) Length() int {
	return s.length
}

func (s ListStack[T]) IsEmpty() bool {
	return s.length == 0 && s.top == nil
}
