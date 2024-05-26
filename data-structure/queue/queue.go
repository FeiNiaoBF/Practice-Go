package queue

import (
	"container/list"
	"fmt"
)

// queue is a FIFO data structure

type Queue struct {
	data *list.List
}

// CircularQueue is a circular queue, using two pointers
type CircularQueue struct {
	data  []any
	front int
	rear  int
	size  int
	cap   int
}

// NewQueue returns a new queue
func (q *Queue) NewQueue() *Queue {
	return &Queue{data: list.New()}
}

// NewCircularQueue returns a new circular queue
func (q *CircularQueue) NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		data:  make([]any, cap),
		front: 0,
		rear:  0,
		size:  0,
		cap:   cap,
	}
}

// Push adds a value to the end of the queue
func (q *Queue) Push(value any) {
	q.data.PushBack(value)
}

// Pop a value to the end of the queue
func (q *Queue) Pop() any {
	if q.isEmpty() {
		return nil
	}
	elem := q.data.Front()
	q.data.Remove(elem)
	return elem.Value
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.data.Len()
}

// isEmpty return true if queue is empty
func (q *Queue) isEmpty() bool {
	return q.data.Len() == 0
}

// Size returns the number of elements in the queue
func (q *CircularQueue) Size() int {
	return q.size
}

// Push adds a value to the end of the queue
func (q *CircularQueue) Push(value any) {
}

// isFull return true if queue is full
func (q *CircularQueue) isFull() bool {
	return q.size == q.cap
}

// PushFront adds a value to the end of the queue
func (q *CircularQueue) PushFront(val any) error {
	if q.isFull() {
		return fmt.Errorf("Queue is full")
	}
	q.data[q.rear] = val
	q.rear = (q.rear + 1) % q.cap
	q.size++
	return nil
}

// PopRear a value to the end of the queue
func (q *CircularQueue) PopRear() (any, error) {
	if !q.isFull() && q.front == q.rear && q.size == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	val := q.data[q.front]
	q.size--
	q.front = (q.front + 1) % q.cap
	return val, nil
}
