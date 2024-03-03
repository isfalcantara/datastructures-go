package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListStackNewListStack(t *testing.T) {
	stack := NewListStack[int]()

	expected := ListStack[int]{
		length: 0,
		top:    nil,
	}

	assert.Equal(t, expected, stack)
}

func Test_ListStackPush(t *testing.T) {
	t.Run("push a value to an empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		stack.Push(1)

		expectedTop := &node[int]{value: 1}

		assert.Equal(t, stack.length, 1)
		assert.Equal(t, expectedTop, stack.top)
	})

	t.Run("push a value to a non-empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		stack.Push(1)
		stack.Push(2)

		expectedBottom := &node[int]{value: 1}
		expectedTop := &node[int]{value: 2, previous: expectedBottom}

		assert.Equal(t, stack.length, 2)
		assert.Equal(t, expectedTop, stack.top)
		assert.Equal(t, expectedBottom, stack.top.previous)
	})
}

func Test_ListStackPop(t *testing.T) {
	t.Run("pop a value from a stack", func(t *testing.T) {
		stack := NewListStack[int]()

		stack.Push(1)
		stack.Push(2)
		popped, err := stack.Pop()

		if assert.NoError(t, err) {
			assert.Equal(t, 2, popped)
			assert.Equal(t, 1, stack.length)
			assert.Equal(t, 1, stack.top.value)
		}
	})

	t.Run("pop the last value of a stack", func(t *testing.T) {
		stack := NewListStack[int]()
		stack.Push(1)

		popped, err := stack.Pop()

		if assert.NoError(t, err) {
			assert.Equal(t, 1, popped)
			assert.Equal(t, 0, stack.length)
			assert.Nil(t, stack.top)
		}
	})

	t.Run("pop a value from an empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		popped, err := stack.Pop()

		if assert.ErrorContains(t, err, "stack is empty") {
			assert.Equal(t, 0, popped)
		}
	})
}

func Test_ListStackPeek(t *testing.T) {
	t.Run("take a peek at the top of a stack", func(t *testing.T) {
		stack := NewListStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		value, err := stack.Peek()
		if assert.NoError(t, err) {
			assert.Equal(t, 3, stack.length)
			assert.Equal(t, 3, stack.top.value)
			assert.Equal(t, stack.top.value, value)
		}
	})

	t.Run("error when taking a peek at an empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		value, err := stack.Peek()
		if assert.ErrorContains(t, err, "stack is empty") {
			assert.Equal(t, 0, value)
		}
	})
}

func Test_ListStackLength(t *testing.T) {
	t.Run("the length of an empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		assert.Equal(t, 0, stack.Length())
	})

	t.Run("the length of a non-empty stack", func(t *testing.T) {
		stack := NewListStack[int]()
		stack.Push(1)
		stack.Push(1)
		stack.Push(3)

		assert.Equal(t, 3, stack.Length())
	})
}

func Test_ListStackIsEmpty(t *testing.T) {
	t.Run("an empty stack", func(t *testing.T) {
		stack := NewListStack[int]()

		assert.True(t, stack.IsEmpty())
	})

	t.Run("a non-empty stack", func(t *testing.T) {
		stack := NewListStack[int]()
		stack.Push(1)

		assert.False(t, stack.IsEmpty())
	})
}
