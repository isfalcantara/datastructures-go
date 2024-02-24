package datastructures

import "errors"

type LinkedList[T any] struct {
	first  *node[T]
	last   *node[T]
	length int
}

func NewLinkedList[T any]() LinkedList[T] {
	return LinkedList[T]{}
}

func (l *LinkedList[T]) Append(value T) {
	newNode := node[T]{value: value}

	if l.IsEmpty() {
		l.first = &newNode
		l.last = &newNode
		l.length = 1
	} else {
		l.last.next = &newNode
		l.last = &newNode
		l.length++
	}
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := node[T]{value: value}

	if l.IsEmpty() {
		l.first = &newNode
		l.last = &newNode
		l.length = 1
	} else {
		newNode.next = l.first
		l.first = &newNode
		l.length++
	}
}

func (l *LinkedList[T]) PopFirst() (popped T, err error) {
	if l.IsEmpty() {
		return popped, errors.New("list is empty")
	} else if l.length == 1 {
		popped = l.first.value

		l.first = nil
		l.last = nil
		l.length = 0

		return popped, nil
	}

	popped = l.first.value

	l.first = l.first.next
	l.length--

	return popped, nil
}

func (l *LinkedList[T]) PopLast() (popped T, err error) {
	if l.IsEmpty() {
		return popped, errors.New("list is empty")
	} else if l.length == 1 {
		popped = l.last.value

		l.first = nil
		l.last = nil
		l.length = 0

		return popped, nil
	}

	popped = l.last.value
	penultimateNode := l.findNode(l.length - 2)

	l.last = penultimateNode
	l.length--

	return popped, nil
}

func (l LinkedList[T]) At(index int) (value T, err error) {
	if l.IsEmpty() {
		return value, errors.New("list is empty")
	}

	return l.findNode(index).value, nil
}

func (l *LinkedList[T]) Insert(value T, index int) {
	newNode := node[T]{value: value}
	index = mod(index, l.length)
	
	if l.IsEmpty() {
		l.first = &newNode
		l.last = &newNode
		l.length = 1
	} else if index == 0 {
		l.Prepend(value)
	} else if index == l.length {
		l.Append(value)
	} else {
		previousNode := l.findNode(index - 1)
		newNode.next = previousNode.next
		previousNode.next = &newNode
		l.length++
	}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.length == 0 && l.first == nil && l.last == nil
}

func (l LinkedList[T]) Length() int {
	return l.length
}

func (l LinkedList[T]) findNode(index int) *node[T] {
	index = mod(index, l.length)

	node := l.first
	for i := 0; i != index; i++ {
		node = node.next
	}

	return node
}
