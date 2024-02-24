package datastructures

type node[T any] struct {
	value    T
	previous *node[T]
	next     *node[T]
}
