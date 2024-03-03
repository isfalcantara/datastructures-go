package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewStaticStack(t *testing.T) {
	stack := NewStaticStack[int](1)
	expectedData := make([]int, 1)

	assert.Equal(t, 0, stack.length)
	assert.Equal(t, expectedData, stack.data)
}

func Test_StaticStackPush(t *testing.T) {
	t.Run("push a value", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		stack.Push(1)
		expectedData := []int{1, 0}

		assert.Equal(t, 1, stack.length)
		assert.Equal(t, expectedData, stack.data)
	})

	t.Run("push and fill up the stack", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		stack.Push(1)
		stack.Push(2)

		expectedData := []int{1, 2}

		assert.Equal(t, 2, stack.length)
		assert.Equal(t, expectedData, stack.data)
	})

	t.Run("push a value and resize the stack", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		expectedData := []int{1, 2, 3, 0}

		assert.Equal(t, 3, stack.length)
		assert.Equal(t, expectedData, stack.data)
	})
}

func Test_StaticStackPop(t *testing.T) {
	t.Run("pop a value from the stack", func(t *testing.T) {
		stack := NewStaticStack[int](1)
		stack.Push(1)
		assert.Equal(t, 1, stack.length)

		popped, err := stack.Pop()
		if assert.NoError(t, err) {
			assert.Equal(t, 0, stack.length)
			assert.Equal(t, 1, popped)
		}
	})

	t.Run("error when stack is empty", func(t *testing.T) {
		stack := NewStaticStack[int](1)

		popped, err := stack.Pop()
		if assert.ErrorContains(t, err, "stack is empty") {
			assert.Equal(t, 0, stack.length)
			assert.Equal(t, 0, popped)
		}
	})
}

func Test_StaticStackPeek(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		stack := NewStaticStack[int](1)

		value, err := stack.Peek()

		if assert.ErrorContains(t, err, "stack is empty") {
			assert.Equal(t, 0, value)
			assert.Equal(t, 0, stack.length)
		}
	})

	t.Run("stack has values", func(t *testing.T) {
		stack := NewStaticStack[int](1)
		stack.Push(1)

		value, err := stack.Peek()

		if assert.NoError(t, err) {
			assert.Equal(t, 1, value)
			assert.Equal(t, 1, stack.length)
		}
	})
}

func Test_StaticStackLength(t *testing.T) {
	t.Run("stack is empty", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		assert.Equal(t, 0, stack.Length())
	})

	t.Run("pushed some values", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		stack.Push(1)
		stack.Push(2)

		assert.Equal(t, 2, stack.Length())
	})
}

func Test_StaticStackIsEmpty(t *testing.T) {
	t.Run("stack is empty", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		assert.True(t, stack.IsEmpty())
	})

	t.Run("pushed a value", func(t *testing.T) {
		stack := NewStaticStack[int](2)
		stack.Push(1)

		assert.False(t, stack.IsEmpty())
	})

	t.Run("popped last value", func(t *testing.T) {
		stack := NewStaticStack[int](1)
		stack.Push(1)
		_, err := stack.Pop()
		if assert.NoError(t, err) {
			assert.True(t, stack.IsEmpty())
		}
	})
}
