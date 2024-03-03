package datastructures

import "errors"

type StaticStack[T any] struct {
	length int
	data   []T
}

func NewStaticStack[T any](initialCapacity int) StaticStack[T] {
	return StaticStack[T]{
		data: make([]T, initialCapacity),
	}
}

func (s *StaticStack[T]) Push(value T) {
	if s.length == len(s.data) {
		s.resize()
	}

	s.data[s.length] = value
	s.length++
}

func (s *StaticStack[T]) Pop() (popped T, err error) {
	if s.IsEmpty() {
		return popped, errors.New("stack is empty")
	}

	s.length--
	return s.data[s.length], nil
}

func (s StaticStack[T]) Peek() (value T, err error) {
	if s.IsEmpty() {
		return value, errors.New("stack is empty")
	}

	return s.data[s.length-1], nil
}

func (s StaticStack[T]) Length() int {
	return s.length
}

func (s StaticStack[T]) IsEmpty() bool {
	return s.length < 1
}

func (s *StaticStack[T]) resize() {
	var newCap int

	if len(s.data) > 1024 {
		newCap = (len(s.data) * 125) / 100
	} else {
		newCap = len(s.data) * 2
	}

	newData := make([]T, newCap)
	copy(newData, s.data)
	s.data = newData
}
