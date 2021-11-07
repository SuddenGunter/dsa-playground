package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_IsEmpty(t *testing.T) {
	cases := []struct {
		queue    *Queue
		expected bool
	}{
		{
			queue:    NewQueue(),
			expected: true,
		},
		{
			queue: &Queue{
				head: &Element{
					Data: 99,
					next: nil,
				},
				last: &Element{
					Data: 99,
					next: nil,
				},
			},
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.expected, test.queue.IsEmpty())
	}
}

func TestQueue_Peek(t *testing.T) {
	expectedEl := &Element{
		Data: 99,
		next: nil,
	}

	cases := []struct {
		queue    *Queue
		expected *Element
	}{
		{
			queue:    NewQueue(),
			expected: nil,
		},
		{
			queue: &Queue{
				head: expectedEl,
				last: expectedEl,
			},
			expected: expectedEl,
		},
	}

	for _, test := range cases {
		assert.Equal(t, test.expected, test.queue.Peek())
	}
}

func TestQueue_Enqueue_OnEmptyQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(Element{
		Data: 999,
		next: nil,
	})

	assert.False(t, queue.IsEmpty())
	assert.NotNil(t, queue.head)
	assert.NotNil(t, queue.last)
	assert.Equal(t, queue.head, queue.last)
	assert.Equal(t, 999., queue.Peek().Data)
}

func TestQueue_Enqueue_OnNonEmptyQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(Element{
		Data: 1,
	})
	queue.Enqueue(Element{
		Data: 2,
	})
	queue.Enqueue(Element{
		Data: 3,
	})

	assert.False(t, queue.IsEmpty())
	assert.NotNil(t, queue.head)
	assert.NotNil(t, queue.last)
	assert.Equal(t, 1., queue.Peek().Data)
	assert.Equal(t, 3., queue.last.Data)
}

func TestQueue_Dequeue_OnEmptyQueue(t *testing.T) {
	queue := NewQueue()
	assert.Nil(t, queue.Dequeue())
}

func TestQueue_Dequeue_OnNonEmptyQueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(Element{
		Data: 1,
	})
	queue.Enqueue(Element{
		Data: 2,
	})
	queue.Enqueue(Element{
		Data: 3,
	})

	first, second, last := queue.Dequeue(), queue.Dequeue(), queue.Dequeue()

	assert.True(t, queue.IsEmpty())
	assert.Nil(t, queue.head)
	assert.Nil(t, queue.last)
	assert.Equal(t, 1., first.Data)
	assert.Equal(t, 2., second.Data)
	assert.Equal(t, 3., last.Data)
}
