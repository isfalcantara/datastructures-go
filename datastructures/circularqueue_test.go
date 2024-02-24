package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewCircularQueue(t *testing.T) {
	expected := CircularQueue[int]{
		first: -1,
		last:  -1,
		data:  make([]int, 1),
	}

	queue := NewCircularQueue[int](1)

	assert.Equal(t, expected, queue)
}

func Test_CircularQueueEnqueue(t *testing.T) {
	t.Run("enqueue a value", func(t *testing.T) {
		queue := NewCircularQueue[int](1)

		queue.Enqueue(1)

		expected := CircularQueue[int]{
			first: 0,
			last:  0,
			data:  []int{1},
		}

		assert.Equal(t, expected, queue)
	})

	t.Run("enqueue a value and wrap around", func(t *testing.T) {
		queue := NewCircularQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()
		queue.Enqueue(4)

		expected := CircularQueue[int]{
			first: 1,
			last:  0,
			data:  []int{4, 2, 3},
		}

		assert.Equal(t, expected, queue)
	})

	t.Run("equeue a value and resize the queue", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)

		expected := CircularQueue[int]{
			first: 0,
			last:  2,
			data:  []int{1, 2, 3, 0},
		}

		assert.Equal(t, expected, queue)
	})

	t.Run("equeue a value in the middle and resize the queue", func(t *testing.T) {
		queue := NewCircularQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()
		queue.Enqueue(4)
		queue.Enqueue(5)

		expected := CircularQueue[int]{
			first: 0,
			last:  3,
			data:  []int{2, 3, 4, 5, 0, 0},
		}

		assert.Equal(t, expected, queue)
	})
}

func Test_CircularQueueDequeue(t *testing.T) {
	t.Run("dequeue a value", func(t *testing.T) {
		queue := NewCircularQueue[int](2)

		queue.Enqueue(1)
		queue.Enqueue(2)

		dequeued, err := queue.Dequeue()

		if assert.NoError(t, err) {
			assert.Equal(t, 1, queue.first)
			assert.Equal(t, 1, queue.last)
			assert.Equal(t, 1, dequeued)
		}
	})

	t.Run("dequeue last value", func(t *testing.T) {
		queue := NewCircularQueue[int](1)
		queue.Enqueue(1)

		dequeued, err := queue.Dequeue()
		if assert.NoError(t, err) {
			assert.Equal(t, -1, queue.first)
			assert.Equal(t, -1, queue.last)
			assert.Equal(t, 1, dequeued)

		}
	})

	t.Run("error when queue is empty", func(t *testing.T) {
		queue := NewCircularQueue[int](1)

		dequeued, err := queue.Dequeue()
		if assert.ErrorContains(t, err, "queue is empty") {
			assert.Equal(t, 0, dequeued)
		}
	})
}

func Test_CircularQueueFirst(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		queue := NewCircularQueue[int](1)

		value, err := queue.First()
		if assert.ErrorContains(t, err, "queue is empty") {
			assert.Equal(t, 0, value)
		}
	})

	t.Run("first value at first position", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		value, err := queue.First()
		if assert.NoError(t, err) {
			assert.Equal(t, 1, value)
		}
	})

	t.Run("first value at another position", func(t *testing.T) {
		queue := NewCircularQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()

		value, err := queue.First()
		if assert.NoError(t, err) {
			assert.Equal(t, 2, value)
		}
	})
}

func Test_CircularQueueLast(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		queue := NewCircularQueue[int](1)

		value, err := queue.Last()
		if assert.ErrorContains(t, err, "queue is empty") {
			assert.Equal(t, 0, value)
		}
	})

	t.Run("last value at last position", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		value, err := queue.Last()
		if assert.NoError(t, err) {
			assert.Equal(t, 2, value)
		}
	})

	t.Run("last value at another position", func(t *testing.T) {
		queue := NewCircularQueue[int](3)
		queue.Enqueue(1)
		queue.Enqueue(2)
		queue.Enqueue(3)
		queue.Dequeue()
		queue.Enqueue(4)

		value, err := queue.Last()
		if assert.NoError(t, err) {
			assert.Equal(t, 4, value)
		}
	})
}

func Test_CircularQueueLength(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		assert.Equal(t, 0, queue.Length())
	})

	t.Run("enqueued some values", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		assert.Equal(t, 2, queue.Length())
	})

	t.Run("dequeue a value", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		_, err := queue.Dequeue()
		if assert.NoError(t, err) {
			assert.Equal(t, 1, queue.Length())
		}
	})
}

func Test_CircularQueueIsEmpty(t *testing.T) {
	t.Run("queue is empty", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		assert.True(t, queue.IsEmpty())
	})

	t.Run("enqueue some values", func(t *testing.T) {
		queue := NewCircularQueue[int](2)
		queue.Enqueue(1)
		queue.Enqueue(2)

		assert.False(t, queue.IsEmpty())
	})

}
