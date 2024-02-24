package datastructures

import "errors"

type CircularQueue[T any] struct {
	first int
	last  int
	data  []T
}

func NewCircularQueue[T any](initialCapacity int) CircularQueue[T] {
	return CircularQueue[T]{
		first: -1,
		last:  -1,
		data:  make([]T, initialCapacity),
	}
}

func (q *CircularQueue[T]) Enqueue(value T) {
	if q.Length() == q.cap() {
		q.resize()
	}

	if q.IsEmpty() {
		q.first = 0
	}
	newLast := mod(q.last+1, q.cap())
	q.data[newLast] = value
	q.last = newLast
}

func (q *CircularQueue[T]) Dequeue() (dequed T, err error) {
	if q.IsEmpty() {
		return dequed, errors.New("queue is empty")
	}

	dequed = q.data[q.first]
	newFirst := mod(q.first+1, q.cap())

	if newFirst == mod(q.last+1, q.cap()) {
		q.clear()
	} else {
		q.first = newFirst
	}

	return dequed, nil
}

func (q CircularQueue[T]) First() (value T, err error) {
	if q.IsEmpty() {
		return value, errors.New("queue is empty")
	}

	return q.data[q.first], nil
}

func (q CircularQueue[T]) Last() (value T, err error) {
	if q.IsEmpty() {
		return value, errors.New("queue is empty")
	}

	return q.data[q.last], nil
}

func (q CircularQueue[T]) Length() int {
	if q.IsEmpty() {
		return 0
	}

	return mod(q.last-q.first, q.cap()) + 1
}

func (q CircularQueue[T]) IsEmpty() bool {
	return q.first == -1 && q.last == -1
}

func (q *CircularQueue[T]) clear() {
	q.first = -1
	q.last = -1
}

func (q *CircularQueue[T]) resize() {
	newData := make([]T, q.cap()*2)

	i := q.first
	j := 0

	for j < q.cap() {
		newData[j] = q.data[i]

		i = mod(i+1, q.cap())
		j++
	}

	q.first = 0
	q.last = q.cap() - 1
	q.data = newData
}

func (q CircularQueue[T]) cap() int {
	return len(q.data)
}
